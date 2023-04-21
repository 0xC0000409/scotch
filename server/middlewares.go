package server

import (
	"github.com/gin-gonic/gin"
)

func registerMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}
