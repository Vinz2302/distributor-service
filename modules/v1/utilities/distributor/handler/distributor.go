package handler

import (
	"distributor-service/modules/v1/utilities/distributor/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistributorHandler struct{}

func NewDistributorHandler() *DistributorHandler {
	return &DistributorHandler{}
}

func (handler *DistributorHandler) Login(c *gin.Context) {
	var (
		input model.LoginInput
		// res   model.LoginMessage
	)

	err := c.ShouldBindQuery(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// res.Success = "false"
	// res.Reason = "Wrong Username or Password"

	c.JSON(http.StatusOK, "Success")
}

func (handler *DistributorHandler) Submit(c *gin.Context) {
	var (
		// input model.SubmitInput
		res model.SubmitMessage
	)

	res.Success = "True"
	res.EDisc = "DPF2024"

	c.JSON(http.StatusOK, res)
}
