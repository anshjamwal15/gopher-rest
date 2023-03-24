package routes

import (
	"gopher-rest/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Post("/create", controllers.CreateOrganization)

	route.Post("/add", controllers.AddUser)

	route.Delete("/delete/:userid/:orgid", controllers.DeleteUser)

	route.Get("/view/:userid", controllers.ViewUser)

	route.Get("/all/:orgid", controllers.GetUserList)

}
