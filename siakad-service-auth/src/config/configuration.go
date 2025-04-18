package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	AppVersion    string
	AppDebugMode  string
	AppPort       string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	JwtSecret     string
	JWTExpiration uint64
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	log.Println("Loading configuration from environment variables...")

	jwtExpiration, err := strconv.ParseUint(getEnv("JWT_EXPIRATION", "3600"), 10, 16)
	if err != nil {
		log.Println("Error parsing JWT_EXPIRATION, using default value of 3600")
		jwtExpiration = 3600
	}

	AppConfig = &Config{
		AppName:       getEnv("APP_NAME", "siakad-service-auth"),
		AppVersion:    getEnv("APP_VERSION", "1.0.0"),
		AppDebugMode:  getEnv("APP_DEBUG", "true"),
		AppPort:       getEnv("APP_PORT", "8081"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBName:        getEnv("DB_NAME", "siakad"),
		JwtSecret:     getEnv("JWT_SECRET", "defaultsecret"),
		JWTExpiration: jwtExpiration,
	}

	log.Println("Configuration loaded successfully")
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
