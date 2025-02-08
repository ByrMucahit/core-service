package repository

import (
	"core-service/models"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB

func ConnectDatabase() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER_NAME")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	fmt.Println("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, pass)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Database connected successfully")

	err = DB.AutoMigrate(&models.Dress{})
	if err != nil {
		log.Fatalf("Error migrating database schema: %v", err)
	}
}
