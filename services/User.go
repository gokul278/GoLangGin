package services

import (
	"fmt"
	internal "golangwithgin/internal/model"

	"gorm.io/gorm"
)

type Userservices struct {
	db *gorm.DB
}

func (u *Userservices) InitService(database *gorm.DB) {
	u.db = database
	u.db.AutoMigrate(&internal.Userdata{})
}

type User struct {
	Username string
	Password string
}

func (u *Userservices) GetUserService() []User {
	data := []User{
		{Username: "gokul", Password: "12345678"},
		{Username: "raja", Password: "12345678"},
	}
	return data
}

func (u *Userservices) CreateUserService() string {

	err := u.db.Create(&internal.Userdata{
		Username: "Gokul",
		Password: "12345678",
	})

	if err != nil {
		fmt.Print(err)
	}
	return "Successfully Inserted"
}
