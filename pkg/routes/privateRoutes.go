package routes

import (
	"gopher-rest/app/controllers"
	"gopher-rest/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Post("/create", middleware.JWTProtected(), controllers.CreateOrganization)

	route.Post("/add", middleware.JWTProtected(), controllers.AddUser)

	route.Delete("/delete/:userid/:orgid", middleware.JWTProtected(), controllers.DeleteUser)

	route.Get("/view/:userid", middleware.JWTProtected(), controllers.ViewUser)

	route.Get("/all/:orgid", middleware.JWTProtected(), controllers.GetUserList)

}
