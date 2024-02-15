package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET All Openings",
	})
}

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "CREATE Opening",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "UPDATE Opening",
	})
}

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}
