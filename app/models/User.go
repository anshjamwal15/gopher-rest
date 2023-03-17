package models

import (
	"fmt"
	u "gopher-rest/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	// Role        Role
}

// func (user *User) Validate() (map[string]interface{}, bool) {

// 	if len(user.Username) > 0 {
// 		return u.Message(false, "Please enter Username"), false
// 	}

// 	if len(user.Password) < 6 {
// 		return u.Message(false, "Please enter valid Password"), false
// 	}

// 	temp := &User{}

// 	err := GetDB().Table("user").Where("Username = ?", user.Username).First(temp).Error

// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return u.Message(false, "Connection error. Please retry"), false
// 	}

// 	if temp.Username != "" {
// 		return u.Message(false, "Username already in use by another user."), false
// 	}

// 	return u.Message(false, "Requirement passed"), true

// }

func (user *User) Create() map[string]interface{} {

	fmt.Println(user)
	fmt.Println(user.Username)

	// if res, ok := user.Validate(); !ok {
	// 	return res
	// }

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return u.Message(false, "Struct Validation Error.")
	}

	GetDB().Create(user)

	response := u.Message(true, "Account has been created")

	response["user"] = user

	return response
}
