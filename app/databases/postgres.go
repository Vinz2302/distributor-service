package databases

import (
	"log"
	"time"

	"distributor-service/app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(conf *config.Conf) (*gorm.DB, error) {
	var (
		dsn       string
		connValue int
	)

	switch conf.App.Mode {
	case "develop":
		dsn = "host=" + conf.Db_develop.Host + " user=" + conf.Db_develop.User + " password=" + conf.Db_develop.Pass + " dbname=" + conf.Db_develop.Name + " port=" + conf.Db_develop.Port + " sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 50
	case "staging":
		dsn = "host=" + conf.Db_staging.Host + " user=" + conf.Db_staging.User + " password=" + conf.Db_staging.Pass + " dbname=" + conf.Db_staging.Name + " port=" + conf.Db_staging.Port + " sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 300
	case "production":
		dsn = "host=" + conf.Db_prod.Host + " user=" + conf.Db_prod.User + " password=" + conf.Db_prod.Pass + " dbname=" + conf.Db_prod.Name + " port=" + conf.Db_prod.Port + " sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 800
	default:
		dsn = "host=" + conf.Db.Host + " user=" + conf.Db.User + " password=" + conf.Db.Pass + " dbname=" + conf.Db.Name + " port=" + conf.Db.Port + " sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 150
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("database not connected")
	} else {
		log.Println("database connected")
	}

	//set max connection pooling
	conn, errCon := db.DB()
	if errCon != nil {
		log.Fatal("gorm.Connect Pool: ", errCon)
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(connValue)
	conn.SetConnMaxIdleTime(time.Minute)

	return db, err

}
