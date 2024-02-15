package router

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET All Openings",
			})
		})

		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusCreated, gin.H{
				"message": "Opening created!",
			})
		})
	}

	v1.PUT("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Opening updated!",
		})
	})

	v1.DELETE("/opening", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Opening deleted!",
		})
	})
}
