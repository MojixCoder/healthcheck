package config

import "os"

// appConfig is application config
var appConfig *AppConfig

// AppConfig is application config structure
type AppConfig struct {
	PORT string
	MongoURI string
	DBName string
}

// Init creates and sets application config
func Init() {
	loadEnvFile()
	appConfig = &AppConfig{
		PORT: os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
		DBName: os.Getenv("DB_NAME"),
	}
}

// GetAppConfig returns application config
func GetAppConfig() *AppConfig {
	return appConfig
}
