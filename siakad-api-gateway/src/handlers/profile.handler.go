package handlers

import (
	"io"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/config"
	"github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/utils"
)

func GetProfileHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Forward request ke auth service
	targetURL := config.AppConfig.AuthServiceURL + "/profile"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := utils.ForwardRequest("GET", targetURL, headers, body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "auth service unreachable"})
		return
	}

	responseBody, err := utils.ReadBody(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read response"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}

func UpdateProfileHandler(c *gin.Context) {
}

func ChangePasswordHandler(c *gin.Context) {
}

func UploadAvatarHandler(c *gin.Context) {
}