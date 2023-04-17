package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
	}
}

func registerMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(SimpleMiddleware())
}
