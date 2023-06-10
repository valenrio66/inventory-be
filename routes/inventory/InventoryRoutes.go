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

	app.Use(inventoryHandler.Authenticate)
	app.Get("/getme", inventoryHandler.GetMe)
	app.Get("/", inventoryHandler.GetAllParfume)
	app.Get("/:id_parfume", inventoryHandler.GetSingleParfume)
	app.Post("/", inventoryHandler.CreateParfume)
	app.Put("/update/:id_parfume", inventoryHandler.UpdateParfume)
	app.Delete("/delete/:id_user", inventoryHandler.DeleteParfume)
	app.Post("/logout", inventoryHandler.LogoutUser)
}
