package models

import (
	"gopher-rest/pkg/utils"
	u "gopher-rest/pkg/utils"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id         int `gorm:"primaryKey"`
	Username   string
	Password   string
	Created_At time.Time
	Updated_At time.Time
	Role       string
	Token      string `gorm:"-"`
	// `gorm:"ForeignKey:Id;references:Id;belongsTo"`
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

	if check, msg := utils.CreateUserValidator(user.Username, user.Password); check != true {
		return u.Message(false, msg)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	user.Created_At = time.Now()
	user.Updated_At = time.Now()

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return u.Message(false, "Struct Validation Error.")
	}

	if strings.Contains(user.Role, "ROLE_ADMIN") {
		user.Role = "ROLE_ADMIN"
	} else {
		user.Role = "ROLE_USER"
	}

	GetDB().Create(user)

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		u.Message(false, "Error creating account. Please retry")
	}

	user.Token = token
	user.Password = ""

	response := u.Message(true, "Account has been created")

	response["user"] = user

	return response
}

func Login(username, password string) map[string]interface{} {

	user := &User{}

	err := GetDB().Table("users").Where("username = ?", username).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	user.Password = ""

	token, err := utils.GenerateNewAccessToken()
	user.Token = token

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp

}

func FindUserById(id int) *User {

	temp := User{Id: id}

	err := GetDB().First(&temp, "id = ?", id).Error

	if err != nil {
		return &temp
	}
	temp.Password = ""

	return &temp
}
