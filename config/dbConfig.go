package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Yefhem/hello-world-case/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ----------> Connection to DB and gets env fields
func ConnectDB() *gorm.DB {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file!")
	}

	// -------- Get env and assign
	get_username := os.Getenv("DB_USER")
	get_password := os.Getenv("DB_PASSWORD")
	get_DBName := os.Getenv("DB_NAME")
	get_host := os.Getenv("DB_HOST")
	get_port := os.Getenv("DB_PORT")
	get_driver := os.Getenv("DB_DRIVER")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", get_host, get_username, get_password, get_DBName, get_port)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", get_driver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", get_driver)
	}

	DB.AutoMigrate(&models.Ticket{}, &models.User{})

	return DB
}
