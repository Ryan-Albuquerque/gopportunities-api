package handler

import (
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, statusCode int, err error) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func SendSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
