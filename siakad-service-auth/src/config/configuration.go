package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppVersion string
	AppDebugMode    string
	AppPort    string
	JwtSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	log.Println("Loading configuration from environment variables...")

	AppConfig = &Config{
        AppName:    getEnv("APP_NAME", "siakad-service-auth"),
        AppVersion: getEnv("APP_VERSION", "1.0.0"),
        AppDebugMode:    getEnv("APP_DEBUG", "true"),
		AppPort:    getEnv("APP_PORT", "8081"),
		JwtSecret:  getEnv("JWT_SECRET", "defaultsecret"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "siakad"),
	}

	log.Println("Configuration loaded successfully")
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
