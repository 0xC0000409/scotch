package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UriIntId struct {
	ID uint `uri:"id" binding:"required"`
}

func BindJsonPretty[T any](c *gin.Context, form T) error {
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		prettyValidationErrors(c)
		return err
	}
	return nil
}
