package inventoryRoutes

import (
	inventoryHandler "inventory-be/internal/handler/inventory"

	"github.com/gofiber/fiber/v2"
)

func SetupInvRoutes(router fiber.Router) {
	app := router.Group("/inventory")
	// Register a user
	app.Post("/register", inventoryHandler.RegisterUser)
	app.Post("/login", inventoryHandler.LoginUser)
}
