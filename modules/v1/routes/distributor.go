package routes

import (
	"distributor-service/app/config"
	"distributor-service/modules/v1/utilities/distributor/handler"

	"github.com/gin-gonic/gin"
)

func Distributor(router *gin.Engine, conf *config.Conf, distributorHandler handler.DistributorHandler) {
	v1 := router.Group("/v1/distributor")

	v1.POST("/login", distributorHandler.Login)
	v1.POST("/submit", distributorHandler.Submit)
}
