package router

import (
	"github.com/GiovanePS/go-api-study/handler"
	gin "github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	handler.Init()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
