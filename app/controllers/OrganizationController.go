package controllers

import (
	"gopher-rest/app/models"
	"gopher-rest/pkg/payload/request"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {

	org := &request.CreateOrgRequest{}

	if err := c.BodyParser(org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fetchedUser := models.FindById(org.Created_By)

	validate := validator.New()

	if err := validate.StructPartial(org, "Created_by"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if strings.Compare(fetchedUser.Role, "ROLE_ADMIN") != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "You're not authorized user to create org.",
		})
	}

	newOrg := &models.Organization{
		Name:       org.Name,
		Created_By: fetchedUser.Id,
		Created_At: time.Now(),
		Updated_At: time.Now(),
		Users:      []models.User{*fetchedUser},
	}

	resp := newOrg.Create()

	return c.JSON(resp)
}
