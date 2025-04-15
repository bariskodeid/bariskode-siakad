package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    AppPort        string
    AuthServiceURL string
    JwtSecret      string
    JwtExpiration  string
    JwtIssuer      string
}

var AppConfig *Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    AppConfig = &Config{
        AppPort:        getEnv("APP_PORT", "8080"),
        AuthServiceURL: getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
        JwtSecret:      getEnv("JWT_SECRET", "your_jwt_secret"),
        JwtExpiration: getEnv("JWT_EXPIRATION", "1h"),
        JwtIssuer:      getEnv("JWT_ISSUER", "your_jwt_issuer"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
