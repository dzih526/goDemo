package v1

import (
	"github.com/gin-gonic/gin"
	"standardProject/service"
)

func Service1func(c *gin.Context)  {
	// get parameters and call service functions, return response
	service.ServiceFunc()
}