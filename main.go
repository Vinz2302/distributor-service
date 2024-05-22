package main

import (
	middleware "distributor-service/app/middlewares"
	driver "distributor-service/driver"
	"distributor-service/modules/v1/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	conf := driver.Conf
	if driver.ErrConf != nil {
		log.Fatal(driver.ErrConf)
	}

	// db := driver.DB

	router := gin.Default()
	router.RedirectTrailingSlash = false

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.GinBodyLogMiddleware)

	routes.Distributor(router, nil, *driver.DistributorHandler)

	port := &conf.App.Port

	router.Run(":" + *port)

}
