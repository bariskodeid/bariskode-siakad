package utils

import (
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userUUID string, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid": userUUID,
		"email":     email,
		"role":      role,
		"exp":       time.Now().Add(time.Duration(config.AppConfig.JWTExpiration) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JwtSecret))
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.AppConfig.JwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	userUUID := claims["user_uuid"].(string)
	email := claims["email"].(string)
	role := claims["role"].(string)
	return GenerateJWT(userUUID, email, role)
}

func GetUserUUID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.AppConfig.JwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	userUUID := claims["user_uuid"].(string)
	return userUUID, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.AppConfig.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
