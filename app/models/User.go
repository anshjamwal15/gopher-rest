package models

import (
	u "gopher-rest/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Role     Role
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if len(user.Username) > 0 {
		return u.Message(false, "Please enter Username"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Please enter valid Password"), false
	}

	temp := &User{}

	err := GetDB().Table("user").Where("Username = ?", user.Username).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Username != "" {
		return u.Message(false, "Username already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true

}

// func (user *User) CreateUser() map[string]interface{} {

// 	if res, ok := user.Validate(); !ok {
// 		return res
// 	}

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	user.Password = string(hashedPassword)

// }
