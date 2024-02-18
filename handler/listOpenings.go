package handler

import (
	"net/http"

	"github.com/GiovanePS/go-api-study/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error on listing openings")
		return
	}

	sendSuccess(context, "list openings", openings)
}
