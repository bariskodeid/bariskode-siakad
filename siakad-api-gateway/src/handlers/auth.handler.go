package handlers

import (
    "io"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/config"
    "github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/utils"
)

func LoginHandler(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    // Forward request ke auth service
    targetURL := config.AppConfig.AuthServiceURL + "/login"
    headers := map[string]string{
        "Content-Type": "application/json",
    }

    resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
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

func LogoutHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Forward request ke auth service
	targetURL := config.AppConfig.AuthServiceURL + "/logout"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
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

func RegisterHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Forward request ke auth service
	targetURL := config.AppConfig.AuthServiceURL + "/register"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
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