package main

import (
	"gopher-rest/cmd/server"
	"gopher-rest/pkg/configs"
	"gopher-rest/pkg/middleware"
	"gopher-rest/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	routes.PrivateRoutes(app)

	server.StartServer(app)
}
