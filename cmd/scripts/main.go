package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"simple_go_gin_gorm_postgres_be_template/db/migrations"

	"github.com/joho/godotenv"
)

func main() {
	mgrtUp := flag.Bool("mgrt-up", false, "Jalankan database migration up")
	mgrtDn := flag.Bool("mgrt-dn", false, "Jalankan database migration down")
	seed := flag.Bool("seed", false, "Jalankan data seeder (coming soon)")
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: Error loading .env file, continuing with system env")
	}

	if *mgrtUp {
		migrations.MigrateUp()
		os.Exit(0)
	}

	if *mgrtDn {
		migrations.MigrateDown()
		os.Exit(0)
	}

	if *seed {
		fmt.Println("ℹ️  Feature seed terpilih, namun belum diimplementasikan di migrations.")
		os.Exit(0)
	}
	fmt.Println("❌ Error: Silakan masukkan salah satu flag berikut:")
	fmt.Println("   --mgrt-up : Untuk melakukan migration up")
	fmt.Println("   --mgrt-dn : Untuk melakukan migration down")
	fmt.Println("   --seed    : Untuk melakukan data seeding")
	os.Exit(1)
}
