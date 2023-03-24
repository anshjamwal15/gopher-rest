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

// CreateOrg func for create new org.
// @Description Create new org.
// @Summary Create new org
// @Tags Org
// @Accept json
// @Produce json
// @Param data body string true "Data"
// @Success 200 {object} response.CreateOrgResponse
// @Security ApiKeyAuth
// @Router /api/v1/create [post]
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

// AddUserToOrg func for Add User to org.
// @Description Add User to org.
// @Summary Add User
// @Tags Org
// @Accept json
// @Produce json
// @Param data body string true "data"
// @Success 200 {object} models.Organization
// @Security ApiKeyAuth
// @Router /api/v1/add [post]
func AddUser(c *fiber.Ctx) error {

	req := &request.CreateUserRequest{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if check, msg := utils.CreateUserValidator(req.Username, req.Password); !check {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   msg,
		})
	}

	fetchedAdmin := models.FindUserById(req.AdminId)

	validate := validator.New()

	if err := validate.StructPartial(req, "AdminId"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if strings.Compare(fetchedAdmin.Role, "ROLE_ADMIN") != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "You're not authorized user to create user.",
		})
	}

	fetchedUser := models.CheckExistingUser(req.Username)

	org := models.FindByOrgById(req.OrgId)

	if fetchedUser.Role == "ROLE_USER" {

		models.AddUserInOrg(*org, *fetchedUser)

		// if err != false {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": true,
		// 		"msg":   "Please try again later.",
		// 	})
		// }

		userResp := &response.UserResponse{
			Id:       fetchedUser.Id,
			Username: fetchedUser.Username,
			Role:     fetchedUser.Role,
		}

		orgResp := &response.CreateUserResponse{
			OrgName:      org.Name,
			UserResponse: *userResp,
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "User added in organization.",
			"user":  orgResp,
		})
	}

	newUser := &models.User{Username: req.Username, Password: req.Password}

	newUser.Create()

	models.AddUserInOrg(*org, *newUser)

	// if err != false {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Please try again later.",
	// 	})
	// }

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

// DeleteUser func for Delete User.
// @Description Delete User .
// @Summary Delete User
// @Tags Org
// @Produce json
// @Success 200
// @Param orgid path int true "Org ID"
// @Param userid path int true "User ID"
// @Router /api/v1/delete/{userid}/{orgid} [delete]
func DeleteUser(c *fiber.Ctx) error {

	id, er := strconv.Atoi(c.Params("userid"))

	orgId, err := strconv.Atoi(c.Params("orgid"))

	if err != nil && er != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again.",
		})
	}

	fetchedUser := models.FindUserById(id)

	org := models.FindByOrgById(orgId)

	models.DeleteUser(*org, *fetchedUser)

	// if dbErr != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Please try again later.",
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "User deleted successfully.",
	})
}

// ViewUser func for View User
// @Description View User
// @Summary View User
// @Tags Org
// @Produce json
// @Param userid path int true "User ID"
// @Success 200
// @Router /api/v1/view/{userid} [get]
func ViewUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("userid"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please try again.",
		})
	}

	fetchedUser := models.FindUserById(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"user":  fetchedUser,
	})
}

// GetUserList func for View User List.
// @Description View User list in org.
// @Summary View User
// @Tags Org
// @Produce json
// @Param orgid path int true "Org ID"
// @Success 200
// @Router /api/v1/all/{orgid} [get]
func GetUserList(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("orgid"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": false,
			"msg":   "Please try again.",
		})
	}

	userList := models.GetAllUsersInOrg(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": true,
		"users": userList,
	})

}
