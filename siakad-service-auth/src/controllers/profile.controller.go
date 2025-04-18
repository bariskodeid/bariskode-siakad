package controllers

import (
	"net/http"

	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/config"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/repositories"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userEmail := c.Param("email")
	userRepo := repositories.NewUserRepository(config.DB)
	user, err := userRepo.FindByEmail(userEmail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":     user.Uuid,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":     user.Email,
		"phone":     user.Phone,
		"username":  user.Username,
		"role":      user.Role,
		"status":    user.Status,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}