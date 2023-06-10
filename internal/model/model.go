package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	// "gorm.io/gorm"
)

// User model
type User struct {
	// gorm.Model
	IdUser    uint   `gorm:"primaryKey;column:id_user;autoIncrement" json:"-"`
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

type Parfume struct {
	IdParfume       uint   `gorm:"primaryKey;column:id_parfume;autoIncrement" json:"-"`
	NamaParfume     string `gorm:"column:nama_parfume" json:"nama_parfume"`
	JenisParfume    string `gorm:"column:jenis_parfume" json:"jenis_parfume"`
	Merk            string `gorm:"column:merk" json:"merk"`
	Deskripsi       string `gorm:"column:deskripsi" json:"deskripsi"`
	Harga           int    `gorm:"column:harga" json:"harga"`
	TahunPeluncuran string `gorm:"column:thn_peluncuran" json:"thn_peluncuran"`
	Stok            int    `gorm:"column:stok" json:"stok"`
	Ukuran          string `gorm:"column:ukuran" json:"ukuran"`
}
