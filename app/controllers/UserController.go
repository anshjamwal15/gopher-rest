package controllers

import (
	"fmt"
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

func TestData(c *fiber.Ctx) error {

	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		return err
	}

	fmt.Println(u.Username)

	return c.SendStatus(200)
}

func CheckData(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "hello, from server-side",
	})
}
