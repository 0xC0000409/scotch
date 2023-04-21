package server

import (
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/models"
	"github.com/go-playground/validator/v10"
)

var uniqueEmail validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	return ok && db.Instance().Where("email = ?", email).First(&models.User{}).RowsAffected == 0
}

func validatorsToRegister() map[string]validator.Func {
	return map[string]validator.Func{
		"unique_email": uniqueEmail,
	}
}
