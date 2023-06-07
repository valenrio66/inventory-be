package router

import (
	routes "inventory-be/routes/inventory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	inventoryapi := app.Group("/inv", logger.New())

	// Setup the Node Routes
	routes.SetupInvRoutes(inventoryapi)
}
