package controllers

import (
	"net/http"

	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/models"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/repositories"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo *repositories.UserRepository
}

func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: userRepo}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hashing failed"})
		return
	}

	user := &models.User{
		Uuid:      utils.GenerateUuid(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Username:  input.Username,
		Password:  hashed,
		Role:      "student",
		Status:    "active",
	}
	if err := h.UserRepo.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exists?"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.UserRepo.FindByEmail(input.Email)
	if err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
    // Invalidate the token or perform any necessary logout actions
    userUUID := c.Param("user_uuid")
    if userUUID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "missing user UUID"})
        return
    }

    user , err := h.UserRepo.Logout(userUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not logout"})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
    tokenString := c.Request.Header.Get("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
        return
    }

    token, err := utils.ValidateJWT(tokenString)
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        return
    }

    newToken, err := utils.RefreshToken(tokenString)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not refresh token"})
        return
    }

    userUUID, err := utils.GetUserUUID(tokenString)
    if userUUID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "missing user UUID"})
        return
    }

    h.UserRepo.Update(&models.User{
        Uuid: userUUID,
        AccessToken: newToken,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": newToken})
}
