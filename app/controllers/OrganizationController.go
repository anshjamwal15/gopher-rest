package controllers

import (
	"gopher-rest/app/models"
	"gopher-rest/pkg/payload/request"
	"gopher-rest/pkg/payload/response"
	"gopher-rest/pkg/utils"
	"strconv"
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

	fetchedUser := models.FindUserById(org.Created_By)

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

	newOrg.Create()

	resp := &response.CreateOrgResponse{
		Name:      newOrg.Name,
		CreatedBy: newOrg.Id,
		CreatedAt: newOrg.Created_At,
		UpdatedAt: newOrg.Updated_At,
	}

	return c.JSON(resp)
}

func AddUser(c *fiber.Ctx) error {

	req := &request.CreateUserRequest{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if check, msg := utils.CreateUserValidator(req.Username, req.Password); check != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   msg,
		})
	}

	fetchedUser := models.FindUserById(req.AdminId)

	validate := validator.New()

	if err := validate.StructPartial(req, "AdminId"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if strings.Compare(fetchedUser.Role, "ROLE_ADMIN") != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "You're not authorized user to create user.",
		})
	}
	newUser := &models.User{Username: req.Username, Password: req.Password}

	newUser.Create()

	org := models.FindByOrgById(req.OrgId)

	models.AddUserInOrg(*org, *newUser)

	userResp := &response.UserResponse{
		Id:       newUser.Id,
		Username: newUser.Username,
		Role:     newUser.Role,
	}

	orgResp := &response.CreateUserResponse{
		OrgName:      org.Name,
		UserResponse: *userResp,
	}

	return c.JSON(orgResp)
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("userid"))

	orgId, err := strconv.Atoi(c.Params("orgid"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again.",
		})
	}

	fetchedUser := models.FindUserById(id)

	org := models.FindByOrgById(orgId)

	models.DeleteUser(*org, *fetchedUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "User deleted successfully.",
	})
}
