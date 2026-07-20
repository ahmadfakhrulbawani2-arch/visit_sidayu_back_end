package main

import (
	"bufio" // Tambahkan ini
	"flag"
	"fmt"
	"log"
	"os"
	cln "simple_go_gin_gorm_postgres_be_template/db/clean"
	"simple_go_gin_gorm_postgres_be_template/db/migrations"
	"simple_go_gin_gorm_postgres_be_template/db/seeds"
	"strings" // Tambahkan ini

	"github.com/joho/godotenv"
)

func main() {
	mgrtUp := flag.Bool("mgrt-up", false, "Jalankan database migration up")
	mgrtDn := flag.Bool("mgrt-dn", false, "Jalankan database migration down")
	seed := flag.Bool("seed", false, "Jalankan data seeder (coming soon)")
	clean := flag.Bool("clean", false, "Bersihkan data tiap table pada database")
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
		// --- Double Validation Start ---
		fmt.Print("⚠️  WARNING: This action is IRREVERSIBLE. Are you sure you want to drop the database tables? (y/N): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input")
		}

		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			fmt.Println("❌ Migration down cancelled.")
			os.Exit(0)
		}
		// --- Double Validation End ---

		migrations.MigrateDown()
		fmt.Println("✅ Database tables dropped successfully.")
		os.Exit(0)
	}

	if *seed {
		seeds.Seeder()
		os.Exit(0)
	}

	if *clean {
		// --- Double Validation Start ---
		fmt.Print("⚠️  WARNING: This action is IRREVERSIBLE. Are you sure you want to delete all of the database tables data? (y/N): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input")
		}

		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			fmt.Println("❌ Cleaning database cancelled.")
			os.Exit(0)
		}
		cln.CleanAllTables()
		os.Exit(0)
	}

	fmt.Println("❌ Error: Silakan masukkan salah satu flag berikut:")
	fmt.Println("   --mgrt-up : Untuk melakukan migration up")
	fmt.Println("   --mgrt-dn : Untuk melakukan migration down")
	fmt.Println("   --seed    : Untuk melakukan data seeding")
	os.Exit(1)
}
