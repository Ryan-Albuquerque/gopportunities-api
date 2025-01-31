package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ryan-albuquerque/gopportunities-api/api/http/handler"
)

func initializeOpeningRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
	}
}
