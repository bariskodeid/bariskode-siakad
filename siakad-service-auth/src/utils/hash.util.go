package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func GenerateUuid() string {
	return uuid.New().String()
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
