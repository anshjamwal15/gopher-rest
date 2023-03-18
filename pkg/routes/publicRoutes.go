package routes

import (
	"gopher-rest/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {

	route := a.Group("/api/v1")

	route.Post("/user", controllers.Register)

	route.Post("/login", controllers.Login)

}
