package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for validatorName, function := range validatorsToRegister() {
			v.RegisterValidation(validatorName, function)
		}
	}
}

func Init() {
	router := gin.New()
	registerValidators()
	registerMiddlewares(router)
	registerRoutes(router)

	router.Run(fmt.Sprintf(":%s", "80"))
}
