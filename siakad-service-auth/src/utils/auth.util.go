package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/models"
    "github.com/google/uuid"
)

var userStore = make(map[string]models.User) // username as key

func RegisterUser(username, password string) error {
    if _, exists := userStore[username]; exists {
        return errors.New("username already exists")
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{
        ID:       uuid.NewString(),
        Username: username,
        Password: string(hashed),
    }

    userStore[username] = user
    return nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
    user, exists := userStore[username]
    if !exists {
        return nil, errors.New("user not found")
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    return &user, nil
}

func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.AppConfig.JwtSecret))
}

func GetUserProfile(userID string) (*models.User, error) {
    for _, user := range userStore {
        if user.ID == userID {
            return &user, nil
        }
    }
    return nil, errors.New("user not found")
}
