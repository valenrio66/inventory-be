package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Nama      string `gorm:"column:nama" json:"nama"`
	Username  string `gorm:"uniqueIndex" json:"username"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// JWTClaims struct
type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
}
