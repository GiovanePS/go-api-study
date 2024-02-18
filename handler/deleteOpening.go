package handler

import (
	"fmt"
	"net/http"

	"github.com/GiovanePS/go-api-study/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadGateway, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id %s not found.", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error on deleting opening with id %s.", id))
	}

	sendSuccess(context, "delete", opening)
}
