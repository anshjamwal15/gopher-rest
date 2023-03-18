package controllers

import (
	"gopher-rest/app/models"
	"gopher-rest/pkg/payload/request"

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

	resp := models.FindById(org.Created_By)

	// validate := validator.New()

	// if err := validate.StructPartial(resp, "Id"); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// }

	// user := resp

	return c.JSON(resp)

	// if !strings.Compare(org.) {

	// }

}
