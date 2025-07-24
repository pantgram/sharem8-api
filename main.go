package main

import (
	"log"
	"shareitapi/app/models"
	"shareitapi/app/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {
	dsn := "host=localhost user=gorm password=gorm dbname= port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
 	if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Database connection successful!")

    // Example: auto-migrate models
    db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Subscription{})
	router := routers.SetupRouter()
	router.Run(":5001") // Listen and serve on 0.0.0.0:8080
}