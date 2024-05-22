package driver

import (
	config "distributor-service/app/config"
	// database "distributor-service/app/databases"
	// "/app/rabbit"
	// redis_client "/app/redis"
	// helperDatabases "/pkg/helpers/databases"
)

var (
	/*
	 Define .env configuratian
	*/
	Conf, ErrConf = config.Init()

	/*
	 Database connection
	*/
	// DB, _ = database.Connect(&Conf)
	/*
	 Helper Database
	*/
	// HelperDatabase = helperDatabases.InitHelperDatabase(DB)
	/*
	 RabbitMQ connection
	*/
	// RabbitMQ, _ = rabbit.Rabbitmq_connect(&Conf)

	/*
	 Redis connection
	*/
	// RedisClient = redis_client.InitRedis(&Conf)

	// MongoDB, _ = database.ConnectMongo(&Conf)
)
