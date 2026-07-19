package migrations

import (
	"log"
	cfg "simple_go_gin_gorm_postgres_be_template/internal/config"
	"simple_go_gin_gorm_postgres_be_template/internal/models"

	"github.com/joho/godotenv"
)

var ModelsRegistry = []any{
	&models.Blogs{},
	&models.CultureBlog{},
	&models.Demographies{},
	&models.Galleries{},
	&models.Geographies{},
	&models.Images{},
	&models.IndustriesBlog{},
	&models.Officials{},
	&models.Roles{},
	&models.ShopsAndUmkmsBlog{},
	&models.Superadmins{},
	&models.Timelines{},
}

// create all tables
func MigrateUp() {
	log.Println("ℹ️ Starting database migration (migrate up)")
	godotenv.Load()

	cfg.ConnectDB()
	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate up, database is not connected yet")
	}

	err := DB.AutoMigrate(ModelsRegistry...)
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
	DB := cfg.DB
	if DB == nil {
		log.Fatal("❌ Failed to migrate down, database is not connected yet")
	}

	err := DB.Migrator().DropTable(ModelsRegistry...)
	if err != nil {
		log.Fatal("❌ Failed to drop table: ", err)
	}

	log.Println("✅ Migration down success, database is now clean!")
}
