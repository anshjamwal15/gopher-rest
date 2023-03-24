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

	fetchedUser := models.CheckExistingUser(user.Username)

	if fetchedUser.Role != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "User already exists.",
		})
	}

	resp := user.Create()

	if resp["status"] == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   resp["message"],
		})
	}

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

	if resp["status"] == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   resp["message"],
		})
	}

	return c.JSON(resp)
}
