package inventoryHandler

import (
	"inventory-be/config/database"
	"inventory-be/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser godoc
// @Summary Register User.
// @Description Register User.
// @Tags Authentication
// @Accept application/json
// @Param request body model.User true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.User
// @Router /inv/inventory/register [post]
// Register handler
func RegisterUser(c *fiber.Ctx) error {
	var user model.User

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body Invalid",
		})
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Register",
		})
	}

	// Simpan user ke database
	user.Password = string(hashedPassword)
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil Register!",
	})
}

// LoginUser godoc
// @Summary Login User.
// @Description Login User.
// @Tags Authentication
// @Accept application/json
// @Param request body model.User true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.User
// @Router /inv/inventory/login [post]
// Login handler
func LoginUser(c *fiber.Ctx) error {
	var user model.User
	inputPassword := c.FormValue("password")

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body nya Invalid",
		})
	}

	// Cari user di database berdasarkan username
	result := database.DB.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Terjadi Kesalahan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau Password Salah",
		})
	}

	// Buat JWT token
	claims := &model.JWTClaims{
		IdUser: user.IdUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Mendapatkan Token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

// GetMe godoc
// @Summary Data user berdasarkan ID.
// @Description get data user.
// @Tags Authentication
// @Accept application/json
// @Produce json
// @Success 200 {object} model.User
// @Router /inv/inventory/getme/ [get]
// GetMe handler
func GetMe(c *fiber.Ctx) error {
	// Mendapatkan data user yang sedang login melalui JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTClaims)

	// Cari user di database berdasarkan user ID
	var userData model.User
	result := database.DB.First(&userData, claims.IdUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User Tidak Ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk Mendapatkan Data User",
		})
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}

// Middleware untuk otentikasi JWT
func Authenticate(c *fiber.Ctx) error {
	// Mendapatkan token dari header Authorization
	authHeader := c.Get("Authorization")
	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Header Otorisasi Salah",
		})
	}

	// Verifikasi token
	claims := new(model.JWTClaims)
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token Invalid",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Gagal Mengautentikasi Token",
		})
	}

	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token Salah",
		})
	}

	// Menyimpan data user ke local context
	c.Locals("user", tkn)

	return c.Next()
}

func LogoutUser(c *fiber.Ctx) error {
	// Hapus token dari Authorization header
	c.Set("Authorization", "")

	return c.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}
