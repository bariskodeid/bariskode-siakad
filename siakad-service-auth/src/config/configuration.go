package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    AppPort   string
    JwtSecret string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load()
	if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    AppConfig = &Config{
        AppPort:   getEnv("APP_PORT", "8081"),
        JwtSecret: getEnv("JWT_SECRET", "defaultsecret"),
    }
}

func getEnv(key, fallback string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return fallback
}
