package seeds

import (
	"log"
	cfg "visit-sidayu-backend/internal/config"

	"github.com/joho/godotenv"
)

func Seeder() {
	log.Println("ℹ️ Starting database seeding")
	godotenv.Load()

	cfg.ConnectDB()
	defer cfg.DisconnectDB()

	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to seeding, database is not connected yet")
	}

	for _, seed := range All() {
		log.Printf("🌱 Running seed: %s", seed.Name)
		if err := seed.Run(DB); err != nil {
			log.Fatalf("❌ Running seed '%s' failed with error: %v", seed.Name, err)
		}
	}
	log.Println("✅ All seeds completed successfully")
}
