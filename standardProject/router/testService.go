package router

import (
	"github.com/gin-gonic/gin"
	v1 "standardProject/api/v1"
)

func InitTestSeviceRouter(Router *gin.RouterGroup){
	testSeviceRouter := Router.Group("testservice")
	{
		testSeviceRouter.POST("handleRequest",v1.Service1func)
	}
}
