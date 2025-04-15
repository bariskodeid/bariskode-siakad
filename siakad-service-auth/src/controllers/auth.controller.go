package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/utils"
)

type AuthPayload struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Register(c *gin.Context) {
    var payload AuthPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    err := utils.RegisterUser(payload.Username, payload.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func Login(c *gin.Context) {
    var payload AuthPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    user, err := utils.AuthenticateUser(payload.Username, payload.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
