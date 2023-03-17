package models

import (
	u "gopher-rest/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	id        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Role        Role
}

func (user *User) Validate() (map[string]interface{}, bool) {

	temp := &User{}

	err := GetDB().Table("users").Where("Username = ?", user.Username).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Username != "" {
		return u.Message(false, "Username already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true

}

func (user *User) Create() map[string]interface{} {

	if len(user.Username) < 0 {
		return u.Message(false, "Please enter Username")
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Please enter valid Password")
	}

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
