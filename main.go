package main

import (
	"fmt"
	controllers "golangwithgin/controllers"
	internal "golangwithgin/internal/database"
	"golangwithgin/services"

	// services "golangwithgin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := internal.InitDB()

	if db == nil {
		fmt.Print("Error at DB Connection")
	}

	userServices := &services.Userservices{}
	userServices.InitService(db)

	usercontroller := &controllers.Usercontroller{}
	usercontroller.InitUsercontrollerRoutes(r, *userServices)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong from gin",
	// 	})
	// })

	// r.POST("/getRes", func(c *gin.Context) {

	// 	type getResRequest struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password" binding:"required"`
	// 	}

	// 	var request getResRequest

	// 	if err := c.BindJSON(&request); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": "error",
	// 			"error":  "Something went wrong, Try Again" + err.Error(),
	// 		})

	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":   "Status OK (200)",
	// 		"email":    request.Email,
	// 		"password": request.Password,
	// 	})

	// })

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
