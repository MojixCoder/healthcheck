package config

import (
	"github.com/joho/godotenv"
	"log"
)

// loadEnvFile Loads .env file
func loadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
