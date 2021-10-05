package router

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine{
	var Router = gin.Default()
	PrivateGroup := Router.Group("")
	{
		InitTestSeviceRouter(PrivateGroup)
	}
	return Router
}
