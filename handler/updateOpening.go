package handler

import (
	"net/http"

	"github.com/GiovanePS/go-api-study/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error on updating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error on updating opening")
		return
	}

	sendSuccess(context, "update opening", opening)
}
