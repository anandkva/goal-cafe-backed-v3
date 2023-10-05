package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	statusCreated             = http.StatusCreated
	statusBadRequest          = http.StatusBadRequest
	statusUnauthorized        = http.StatusUnauthorized
	statusInternalServerError = http.StatusInternalServerError
	statusConflict            = http.StatusConflict
)

func HandleUnauthorizedError(ctx *gin.Context, message string) {
	ctx.JSON(statusUnauthorized, gin.H{"error": message})
}

func HandleBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(statusBadRequest, gin.H{"error": err.Error()})
}

func HandleInternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(statusInternalServerError, gin.H{"error": message})
}

func HandleConflict(ctx *gin.Context, message string) {
	ctx.JSON(statusConflict, gin.H{"error": message})
}
