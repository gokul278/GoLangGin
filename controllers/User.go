package controller

import (
	internal "golangwithgin/internal/model"
	services "golangwithgin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	Userservices services.Userservices
}

func (u *Usercontroller) InitUsercontrollerRoutes(router *gin.Engine, userService services.Userservices) {
	users := router.Group("/users")
	users.POST("/", u.PostUsers())
	users.POST("/newuser", u.PostNewUser())
	users.PATCH("/updateUser", u.UpdateUser())
	u.Userservices = userService
}

func (u *Usercontroller) PostUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqVal internal.PostUsers

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Something went wrong, Try Again" + err.Error(),
			})

			return
		}

		users, message := u.Userservices.PostUserService(reqVal)
		c.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Way":     "Get Users",
			"data":    users,
			"message": message,
		})
	}
}

func (u *Usercontroller) PostNewUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		var reqVal internal.CreateUserRequest

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Something went wrong, Try Again" + err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Status": "Success",
			"Way":    "Create User",
			"data":   u.Userservices.CreateUserService(reqVal),
		})
	}
}

func (u *Usercontroller) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVal internal.UpdateUserRequest

		if err := c.BindJSON(&reqVal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Something went wrong, Try Again" + err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": u.Userservices.UpdateUserService(reqVal),
		})
	}
}
