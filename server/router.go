package server

import (
	"github.com/0xC0000409/scotch/controllers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/users", controllers.UsersGetController)
		v1.POST("/user", controllers.UserPostController)
	}
}
