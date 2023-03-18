package models

import (
	u "gopher-rest/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type Organization struct {
	Id         int
	Name       string
	Created_by *User
	Created_at time.Time
	Updated_at time.Time
}

func (org *Organization) Create() map[string]interface{} {

	validate := validator.New()

	if err := validate.Struct(org); err != nil {
		return u.Message(false, "Struct Validation Error.")
	}

	GetDB().Create(org)

	response := u.Message(false, "Failed to created organization. Please retry.")

	response["organization"] = org

	return response
}
