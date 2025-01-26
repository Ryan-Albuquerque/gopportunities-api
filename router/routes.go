package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ryan-albuquerque/gopportunities-api/handler"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ListOpenings)
	}
}