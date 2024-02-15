package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "CREATE Opening",
	})

	request := struct {
		role string
	}{}

	context.BindJSON(&request)
}
