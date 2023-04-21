package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UriIntId struct {
	ID int `uri:"id" binding:"required"`
}

func MustBindUri[T any](uriForm T) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindUri(&uriForm); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Next()
	}
}

func MustBindJson[T any](form T) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&form); err != nil {
			return
		}
		c.Next()
	}
}

func MustBindJsonPretty[T any](form T) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(form)
		if err := c.ShouldBindJSON(&form); err != nil {
			c.Error(err).SetType(gin.ErrorTypeBind)
			prettyValidationErrors(c)
			return
		}
		c.Next()
	}
}
