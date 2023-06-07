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
			"message": "Invalid request body",
		})
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	// Simpan user ke database
	user.Password = string(hashedPassword)
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
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

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Cari user di database berdasarkan username
	result := database.DB.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid username or password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	// Buat JWT token
	claims := &model.JWTClaims{
		IdUser: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to login",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

// GetMe handler
func getMe(c *fiber.Ctx) error {
	// Mendapatkan data user yang sedang login melalui JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTClaims)

	// Cari user di database berdasarkan user ID
	var userData model.User
	result := database.DB.First(&userData, claims.IdUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user data",
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
			"message": "Invalid authorization header",
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
				"message": "Invalid token signature",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to authenticate token",
		})
	}

	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Menyimpan data user ke local context
	c.Locals("user", tkn)

	return c.Next()
}
