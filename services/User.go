package services

import (
	"fmt"
	internal "golangwithgin/internal/model"
	"time"

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

func (u *Userservices) PostUserService(reqVal internal.PostUsers) ([]internal.Userdata, string) {
	var users []internal.Userdata

	// var results []struct {
	// 	internal.Userdata
	// 	ProfileName string `gorm:"column:profile_name"`
	// 	// Add other fields from joined table as needed
	// }

	if len(reqVal.ID) > 0 {
		// result := u.db.First(&users, reqVal.ID)
		result := u.db.Raw(`SELECT * FROM public."userdata" WHERE id=?`, reqVal.ID).Scan(&users)

		if result.Error != nil {
			fmt.Println("Get Error:", result.Error)
			return nil, "Failed to fetch users: " + result.Error.Error()
		}
	} else {
		result := u.db.Order("id asc").Find(&users)

		if result.Error != nil {
			fmt.Println("Get Error:", result.Error)
			return nil, "Failed to fetch users: " + result.Error.Error()
		}
	}

	return users, "Successfully retrieved users"
}

func (u *Userservices) CreateUserService(req internal.CreateUserRequest) string {
	user := internal.Userdata{
		Username:  req.Email,
		Password:  req.Password,
		CreatedBy: "Admin",
	}

	dbResult := u.db.Create(&user)
	if dbResult.Error != nil {
		fmt.Println("Insert Error:", dbResult.Error)
		return "Insert Failed"
	}

	return "Successfully Inserted"
}

func (u *Userservices) UpdateUserService(req internal.UpdateUserRequest) string {
	updates := map[string]interface{}{
		"username":   req.Email, // Note: Consider if you really want to set username=email
		"password":   req.Password,
		"updated_by": "Admin",
		"updated_at": time.Now(), // Make sure to update the timestamp
	}

	dbResult := u.db.Model(&internal.Userdata{}).Where("id=?", req.ID).Updates(updates)
	if dbResult.Error != nil {
		fmt.Println("Insert Error:", dbResult.Error)
		return "Insert Failed"
	}

	return "Successfully Updated"
}
