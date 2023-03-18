package controllers

import (
	"gopher-rest/app/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	resp := user.Create()

	return c.JSON(resp)
}

func Login(c *fiber.Ctx) error {

	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	resp := models.Login(user.Username, user.Password)

	return c.JSON(resp)
}
