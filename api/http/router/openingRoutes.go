package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ryan-albuquerque/gopportunities-api/api/http/handler/opening"
)

func initializeOpeningRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ListOpeningsHandler)
	}
}