package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// getEnvDir returns the .env file path
//
// Because tests are in a sub-directory then we need to set .env path to ../.env
func getEnvDir() string {
	dir, _ := os.Getwd()
	envPath := ".env"
	if filepath.Base(dir) == "test" {
		envPath = "../.env"
	}
	return envPath
}

// loadEnvFile Loads .env file
func loadEnvFile() {
	err := godotenv.Load(getEnvDir())
	if err != nil {
		log.Fatal(err)
	}
}
