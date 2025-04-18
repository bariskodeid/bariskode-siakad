package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithSuccess(c *gin.Context, status int, message string, data interface{}) {
    c.JSON(status, APIResponse{
        Status:  status,
        Message: message,
        Data:    data,
    })
}

func RespondWithError(c *gin.Context, status int, message string, err error) {
	if err == nil {
		err = gin.Error{Err: nil, Type: gin.ErrorTypePublic, Meta: nil}
	}
    c.JSON(status, APIResponse{
        Status:  status,
        Message: message,
        Error:   err.Error(),
    })
}