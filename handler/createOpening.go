package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	err := db.Create(&request).Error
	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}
}
