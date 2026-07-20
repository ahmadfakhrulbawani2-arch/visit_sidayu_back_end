package clean

import (
	"log"
	"simple_go_gin_gorm_postgres_be_template/db"
	"slices"

	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"

	"github.com/joho/godotenv"
)

func CleanAllTables() {
	log.Println("ℹ️ Starting database cleaning")
	godotenv.Load()

	cfg.ConnectDB()
	defer cfg.DisconnectDB()

	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to cleaning, database is not connected yet")
	}

	for _, table := range slices.Backward(db.ModelsRegistry) {
		if err := DB.Unscoped().Delete(table, "1=1").Error; err != nil {
			log.Fatalf("❌ Failed to cleaning, cleaning aborted with error: %v", err)
		}
	}
	log.Println("✅ All tables data has been deleted successfully")
}
