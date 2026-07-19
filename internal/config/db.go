package config

import (
	"log"
	"os"

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

	DB = database
	log.Println("🎉 Success to connect database")
}
