package controllers

import (
	"github.com/0xC0000409/scotch/db"
	"github.com/0xC0000409/scotch/forms"
	"github.com/0xC0000409/scotch/models"
	"github.com/0xC0000409/scotch/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func usersGetController(c *gin.Context) {
	var users []models.User
	db.Instance().Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func userGetController(c *gin.Context) {
	var uri pkg.UriIntId
	_ = c.BindUri(&uri)

	c.JSON(http.StatusOK, gin.H{"foo": uri.ID})
}

func userPostController(c *gin.Context) {
	var newUser forms.UserCreateForm
	_ = c.BindJSON(&newUser)

	db.Instance().Create(&models.User{
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
	})

	c.JSON(http.StatusCreated, newUser)
}

func userPatchController(c *gin.Context) {
	var uri pkg.UriIntId
	_ = c.BindUri(&uri)

	var updateUser forms.UserUpdateForm
	_ = c.BindJSON(&updateUser)

	var user models.User
	result := db.Instance().First(&user, uri.ID)

	if result.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Instance().Model(&user).Updates(updateUser)
	c.JSON(http.StatusAccepted, user)
}

func userDeleteController(c *gin.Context) {
	var uri pkg.UriIntId
	_ = c.BindUri(&uri)

	if db.Instance().Delete(&models.User{}, uri.ID).RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/users", usersGetController)
	user := router.Group("user")
	{
		user.GET("/:id", pkg.MustBindUri(pkg.UriIntId{}), userGetController)
		user.POST("/", pkg.MustBindJsonPretty(forms.UserCreateForm{}), userPostController)
		user.Use().PATCH("/:id",
			pkg.MustBindUri(pkg.UriIntId{}),
			pkg.MustBindJsonPretty(forms.UserUpdateForm{}),
			userPatchController,
		)
		user.DELETE("/:id", pkg.MustBindUri(pkg.UriIntId{}), userDeleteController)
	}
}
