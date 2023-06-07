package database

import (
	"fmt"
	"log"
	"strconv"

	"inventory-be/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Mendeklarasikan variabel untuk database
var DB *gorm.DB

// Fungsi ConnectDB() untuk menghubungkan ke database
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Fatal("Error parsing DB_PORT")
	}

	// URL koneksi untuk menghubungkan ke database MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))

	// Menghubungkan ke DB dan menginisialisasikan variabel DB
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connection Opened to Database")
}
