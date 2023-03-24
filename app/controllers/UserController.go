package controllers

import (
	"gopher-rest/app/models"

	"github.com/gofiber/fiber/v2"
)

// CreateUser func for create new user.
// @Description Register new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [post]
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

// Login func for Login new user.
// @Description Log In.
// @Summary create new JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/login [post]
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
