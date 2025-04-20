package handlers

import (
    "io"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/config"
    "github.com/bariskodeid/bariskode-siakad/siakad-api-gateway/src/utils"
)

func LoginHandler(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid request", err)
        return
    }

    // Forward request ke auth service
    targetURL := config.AppConfig.AuthServiceURL + "/login"
    headers := map[string]string{
        "Content-Type": "application/json",
    }

    resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
    if err != nil {
        utils.RespondWithError(c, http.StatusBadGateway, "Auth service unreachable", err)
        return
    }

    responseBody, err := utils.ReadBody(resp)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to read response", err)
        return
    }

    c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}

func LogoutHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Forward request ke auth service
	targetURL := config.AppConfig.AuthServiceURL + "/logout"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Auth service unreachable", err)
		return
	}

	responseBody, err := utils.ReadBody(resp)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to read response", err)
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}

func RegisterHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Forward request ke auth service
	targetURL := config.AppConfig.AuthServiceURL + "/register"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := utils.ForwardRequest("POST", targetURL, headers, body)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Auth service unreachable", err)
		return
	}

	responseBody, err := utils.ReadBody(resp)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to read response", err)
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), responseBody)
}