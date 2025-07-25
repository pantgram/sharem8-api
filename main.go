package main

import (
	"log"
	"sharem8-api/app/config"
	"sharem8-api/app/routers"

	"github.com/joho/godotenv"
)



func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
    log.Fatal("Error loading .env file")
  }
	config.Init()
	router := routers.SetupRouter()
	router.Run(":5001") // Listen and serve on 0.0.0.0:5001
}