package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/utils"
)

type ProfilePayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetProfile(c *gin.Context) {
	userID := c.Param("user_id")
	profile, err := utils.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}