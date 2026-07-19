package migrations

import (
	"log"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	modelConst "simple_go_gin_gorm_postgres_be_template/internal/models/constants"

	"github.com/joho/godotenv"
)

func MigrateUp() {
	log.Println("ℹ️ Starting database migration (migrate up)")
	godotenv.Load()

	cfg.ConnectDB()
	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate up, database is not connected yet")
	}

	err := DB.AutoMigrate(modelConst.ModelsRegistry...)
	if err != nil {
		log.Fatal("❌ Failed to do database migration : ", err)
	}

	log.Println("✅ Migration up success, database is up to date!")
}

func MigrateDown() {
	log.Println("ℹ️ Starting database reset (migrate down)")
	godotenv.Load()
	cfg.ConnectDB()
	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate down, database is not connected yet")
	}

	err := DB.Migrator().DropTable(modelConst.ModelsRegistry...)
	if err != nil {
		log.Fatal("❌ Failed to drop table: ", err)
	}

	log.Println("✅ Migration down success, database is now clean!")
}
