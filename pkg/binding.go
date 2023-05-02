package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UriIntId struct {
	ID uint `uri:"id" binding:"required"`
}

func BindJsonPretty[T any](c *gin.Context, form T) error {
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
		} else {
			c.Error(err).SetType(gin.ErrorTypeBind)
			prettyValidationErrors(c)
		}
		return err
	}
	return nil
}
