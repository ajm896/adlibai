package db

import (
	"fmt"
	"os"

	"github.com/ajm896/adlibai/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects to the database

// ConnectDB connects to the database
func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}
