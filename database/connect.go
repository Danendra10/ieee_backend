package database

import (
	"fmt"
	"log"
	"os"

	"github.com/danendra10/ieee_backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" // Import the PostgreSQL driver
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Update the DSN variable to use the PostgreSQL environment variables
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = database

	err = DB.Migrator().DropTable(&models.AirPolution{})
	if err != nil {
		log.Fatal("Failed to drop column")
	}
	tables := []interface{}{
		&models.AirPolution{},
	}

	for _, table := range tables {
		result := database.AutoMigrate(table)
		if result == nil {
			fmt.Printf("Successfully migrated table: %T\n", table)
		}
	}
}

// package database

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/danendra10/ieee_backend/models"
// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dsn := os.Getenv("DSN")
// 	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database")
// 	}

// 	DB = database

// 	err = DB.Migrator().DropTable(&models.AirPolution{})
// 	if err != nil {
// 		log.Fatal("Failed to drop column")
// 	}
// 	tables := []interface{}{
// 		&models.AirPolution{},
// 	}

// 	for _, table := range tables {
// 		result := database.AutoMigrate(table)
// 		if result == nil {
// 			fmt.Printf("Successfully migrated table: %T\n", table)
// 		}
// 	}
// }
