package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.New()
	registerValidators()
	registerMiddlewares(router)
	registerRoutes(router)

	router.Run(fmt.Sprintf(":%s", "80"))
}
