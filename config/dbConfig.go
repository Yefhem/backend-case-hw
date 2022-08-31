package config

import (
	"github.com/Yefhem/hello-world-case/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ---------->
func ConnectDB() *gorm.DB {

	// if err := godotenv.Load(".env", "test.env"); err != nil {
	// 	log.Fatal("Error loading .env file!")
	// }

	// // -------- Get env and assign
	// get_username := os.Getenv("DB_USER")
	// get_password := os.Getenv("DB_PASSWORD")
	// get_DBName := os.Getenv("DB_NAME")
	// get_host := os.Getenv("DB_HOST")
	// get_port := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", get_host, get_username, get_password, get_DBName, get_port)

	dsn := "host=postgresdb user=spuser password=SPuser96 dbname=project port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Ticket{}, &models.User{})

	DB = db

	return DB
}
