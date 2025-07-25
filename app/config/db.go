package config

import (
	"log"
	"os"

	"sharem8-api/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {

	dsn := os.Getenv("DB_DSN")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 	if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	    log.Println("Database connection successful!")
	


	// Auto-migrate models (for MVP)
	err = DB.AutoMigrate(
		&models.User{},
		&models.Subscription{},
		&models.Group{},
	)

	if err != nil {
		log.Fatal("❌ Failed to migrate database schema:", err)
	}
}
