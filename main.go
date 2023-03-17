package main

import (
	"gopher-rest/pkg/configs"
	"gopher-rest/pkg/middleware"
	"gopher-rest/pkg/routes"
	"gopher-rest/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	utils.StartServer(app)
}
