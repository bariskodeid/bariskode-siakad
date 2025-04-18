package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
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
		utils.RespondWithError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not hash password", err)
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
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not create user", err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "Registered", user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid input", err)
		return
	}

	user, err := h.UserRepo.FindByEmail(input.Email)
	if err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		user.LastFailedLogin = sql.NullTime{Time: time.Now(), Valid: true}
		user.LoginAttempt += 1
		if err := h.UserRepo.Update(user); err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Could not update user", err)
			return
		}
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials", errors.New("invalid credentials"))
		return
	}

	token, err := utils.GenerateJWT(user.Uuid)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not generate token", err)
		return
	}
	user.AccessToken = token
	user.AccessTokenExpiredAt = sql.NullTime{Time: time.Now().Add(time.Duration(config.AppConfig.JWTExpiration) * time.Second), Valid: true}
	user.Status = "active"
	user.LoginAttempt += 1
	user.LastSuccessfulLogin = sql.NullTime{Time: time.Now(), Valid: true}
	if err := h.UserRepo.Update(user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not update user", err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Logged in", gin.H{
		"token": token,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var input struct {
		UserUUID string `json:"user_uuid"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid input", err)
		return
	}
	userUUID := input.UserUUID

	if userUUID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing user UUID", errors.New("missing user UUID"))
		return
	}

	user, err := h.UserRepo.Logout(userUUID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not logout", err)
		return
	}
	if user == nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found", errors.New("user not found"))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Logged out", gin.H{})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing token", errors.New("missing token"))
		return
	}

	token, err := utils.ValidateJWT(tokenString)
	if err != nil || !token.Valid {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	newToken, err := utils.RefreshToken(tokenString)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not refresh token", err)
		return
	}

	userUUID, err := utils.GetUserUUID(tokenString)
	if userUUID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Missing user UUID", errors.New("missing user UUID"))
		return
	}

	h.UserRepo.Update(&models.User{
		Uuid:        userUUID,
		AccessToken: newToken,
	})
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Could not update user", err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Token refreshed", gin.H{
		"token": newToken,
	})
}
