package server

import (
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/models"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var uniqueEmail validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	return ok && db.Instance().First(&models.User{Email: email}).RowsAffected == 0
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique_email", uniqueEmail)
	}
}
