package main

import (
	"gopher-rest/cmd/server"
	_ "gopher-rest/docs"
	"gopher-rest/pkg/configs"
	"gopher-rest/pkg/middleware"
	"gopher-rest/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name Ansh Jamwal
// @contact.email anshjaamwal2002@mail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	routes.PrivateRoutes(app)

	routes.SwaggerRoute(app)

	server.StartServer(app)
}
