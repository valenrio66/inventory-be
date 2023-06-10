package main

import (
	"inventory-be/config/database"
	_ "inventory-be/docs"
	"inventory-be/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Inventory Parfume API
// @description API For Inventory Parfume
// @version 1
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @security ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Connect to the Database
	database.ConnectDB()

	// Start a new fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup the router
	router.SetupRoutes(app)

	// Generate Swagger documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	// Listen on PORT 3000
	app.Listen(":3000")
}
