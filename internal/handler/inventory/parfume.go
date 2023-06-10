package inventoryHandler

import (
	"inventory-be/config/database"
	"inventory-be/internal/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllParfume godoc
// @Summary Get all parfume
// @Description Menampilkan semua data parfume
// @Tags Parfume
// @Accept application/json
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} model.Parfume
// @Router /inv/inventory [get]
func GetAllParfume(c *fiber.Ctx) error {
	var wangi []model.Parfume

	// Find all users in database
	result := database.DB.Find(&wangi)

	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return list of users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Parfume Berhasil Ditampilkan!",
		"data":    wangi,
	})
}

// CreateParfume godoc
// @Summary insert data parfume.
// @Description get data parfume.
// @Tags Parfume
// @Accept application/json
// @Param request body model.Parfume true "Payload Body [RAW]"
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} model.Parfume
// @Router /inv/inventory [post]
func CreateParfume(c *fiber.Ctx) error {
	// Parse request body
	var wangi model.Parfume
	if err := c.BodyParser(&wangi); err != nil {
		return err
	}

	// Insert new parfume into database
	result := database.DB.Create(&wangi)

	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data Parfume Berhasil Ditambahkan!",
		"data":    wangi,
	})
}

// GetSingleParfume godoc
// @Summary Data parfume berdasarkan ID.
// @Description get data parfume.
// @Tags Parfume
// @Accept application/json
// @Param id_parfume path string true "Masukan ID Parfume"
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /inv/inventory/{id_parfume} [get]
func GetSingleParfume(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_parfume")

	// Find user by id_user in database
	var wangi model.Parfume
	result := database.DB.First(&wangi, id)

	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Parfume Tidak Ditermukan!",
		})
	}

	// Check for errors during query
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return user
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    wangi,
	})
}

// UpdateParfume godoc
// @Summary update data parfume.
// @Description update data parfume.
// @Tags Parfume
// @Accept application/json
// @Param id_parfume path string true "Masukan ID Parfume"
// @Param request body model.User true "Payload body [RAW]"
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /inv/inventory/update/{id_parfume} [put]
func UpdateParfume(c *fiber.Ctx) error {
	// Get id_parfume parameter from request url
	id := c.Params("id_parume")

	// Find parfume by id_parume in database
	var wangi model.Parfume
	result := database.DB.First(&wangi, id)

	// Check if parfume exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Parfume Tidak Ditemukan",
		})
	}

	// Parse request body
	var newParfume model.Parfume
	if err := c.BodyParser(&newParfume); err != nil {
		return err
	}

	// Update parfume in database
	result = database.DB.Model(&wangi).Updates(newParfume)

	// Check for errors during update
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Parfume Berhasil Diperbarui!",
		"data":    wangi,
	})
}

// DeleteParfume godoc
// @Summary Hapus Data parfume berdasarkan ID.
// @Description delete data parfume.
// @Tags Parfume
// @Accept application/json
// @Param id_parfume path string true "Masukan ID parfume"
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /inv/inventory/delete/{id_parfume} [delete]
func DeleteParfume(c *fiber.Ctx) error {
	// Get id_parfume parameter from request url
	id := c.Params("id_parfume")

	// Find user by id_parfume in database
	var wangi model.Parfume
	result := database.DB.First(&wangi, id)

	// Check if parfume exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Parfume Tidak Ditemukan",
		})
	}

	// Delete parfume from database
	result = database.DB.Delete(&wangi)

	// Check for errors during deletion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data Parfume Berhasil Dihapus!",
		"data":    wangi,
	})
}
