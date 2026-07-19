package config

import (
	"log"
	"os"
	"simple_go_gin_gorm_postgres_be_template/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	log.Println("🚀 Starting database connection")
	dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URI not found in .env")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect the database : ", err)
	}

	err = database.AutoMigrate(&models.Event{}, models.User{}, models.Image{})
	if err != nil {
		log.Fatal("❌ Failed to do database migration : ", err)
	}

	DB = database
	log.Println("🎉 Success to connect database")
}
