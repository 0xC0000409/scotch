package controllers

import (
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/forms"
	"github.com/0xC0000409/scotch/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UsersGetController(c *gin.Context) {
	var users []models.User
	db.Instance().Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UserPostController(c *gin.Context) {
	var newUser forms.UserCreateForm

	if err := c.BindJSON(&newUser); err != nil {
		log.Print(err)
		return
	}

	db.Instance().Create(&models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	})

	c.JSON(http.StatusCreated, newUser)
}
