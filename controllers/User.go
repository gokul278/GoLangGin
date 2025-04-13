package controllers

import (
	services "golangwithgin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	Userservices services.Userservices
}

func (u *Usercontroller) InitUsercontrollerRoutes(router *gin.Engine, userService services.Userservices) {
	users := router.Group("/users")
	users.GET("/", u.GetUsers())
	users.POST("/newuser", u.PostNewUser())
	u.Userservices = userService
}

func (u *Usercontroller) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Success",
			"Way":    "Get Users",
			"data":   u.Userservices.GetUserService(),
		})
	}
}

func (u *Usercontroller) PostNewUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"Status": "Success",
			"Way":    "Create User",
			"data":   u.Userservices.CreateUserService(),
		})
	}
}
