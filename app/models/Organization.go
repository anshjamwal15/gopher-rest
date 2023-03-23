package models

import (
	"errors"
	u "gopher-rest/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type Organization struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Created_By int
	Created_At time.Time
	Updated_At time.Time
	Users      []User `gorm:"many2many:org_users;"`
}

type OrganizationUser struct {
	OrgId     int `gorm:"primaryKey"`
	UserId    int `gorm:"primaryKey"`
	CreatedAt time.Time
}

func (org *Organization) Create() map[string]interface{} {

	validate := validator.New()

	if err := validate.Struct(org); err != nil {
		return u.Message(false, "Struct Validation Error.")
	}

	GetDB().Create(org)

	response := u.Message(true, "Organization created successfully.")

	response["organization"] = org

	return response
}

func FindByOrgById(id int) *Organization {

	temp := Organization{Id: id}

	err := GetDB().First(&temp, "id = ?", id).Error

	if err != nil {
		return &temp
	}

	return &temp

}

func AddUserInOrg(org Organization, user User) error {
	err := GetDB().Model(&org).Association("Users").Append(&user).Error()
	if err != "" {
		return errors.New(err)
	}
	return db.Error
}

func DeleteUser(org Organization, user User) error {
	err := GetDB().Model(&org).Association("Users").Delete(&user).Error()
	if err != "" {
		return errors.New(err)
	}
	return db.Error
}

func GetAllUsersInOrg(id int) []User {

	temp := Organization{Id: id}

	err := GetDB().Model(&temp).Preload("Users").Find(&temp)

	if err != nil {
		return temp.Users
	}
	return temp.Users
}
