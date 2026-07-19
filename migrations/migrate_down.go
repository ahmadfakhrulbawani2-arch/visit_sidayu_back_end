package main

import (
	"log"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	"simple_go_gin_gorm_postgres_be_template/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("ℹ️ Starting database reset (migrate down)")
	godotenv.Load()
	cfg.ConnectDB()
	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate down, database is not connected yet")
	}

	err := DB.Migrator().DropTable(&models.Event{}, &models.Image{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Failed to drop table: ", err)
	}

	log.Println("✅ Migration down success, database is now clean!")

}
