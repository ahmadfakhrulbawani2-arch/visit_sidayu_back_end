package migrations

import (
	"log"
	"visit-sidayu-backend/db"
	cfg "visit-sidayu-backend/internal/config"

	"github.com/joho/godotenv"
)

// create all tables
func MigrateUp() {
	log.Println("ℹ️ Starting database migration (migrate up)")
	godotenv.Load()

	cfg.ConnectDB()
	defer cfg.DisconnectDB()

	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate up, database is not connected yet")
	}

	err := DB.AutoMigrate(db.ModelsRegistry...)
	if err != nil {
		log.Fatal("❌ Failed to do database migration : ", err)
	}

	log.Println("✅ Migration up success, database is up to date!")
}

// drop all tables
func MigrateDown() {
	log.Println("ℹ️ Starting database reset (migrate down)")
	godotenv.Load()

	cfg.ConnectDB()
	defer cfg.DisconnectDB()

	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate down, database is not connected yet")
	}

	err := DB.Migrator().DropTable(db.ModelsRegistry...)
	if err != nil {
		log.Fatal("❌ Failed to drop table: ", err)
	}

	log.Println("✅ Migration down success, database is now clean!")
}
