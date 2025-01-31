package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handlerSchema "github.com/ryan-albuquerque/gopportunities-api/api/http/handler/schemas"
	"github.com/ryan-albuquerque/gopportunities-api/internal/models"
	"github.com/ryan-albuquerque/gopportunities-api/internal/services"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := models.CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if validation := request.ValidateCreateOpeningRequest(); validation != nil {
		SendError(ctx, http.StatusBadRequest, validation)
		return
	}

	createdId, err := services.CreateOpening(request)
	if err != nil && createdId == 0 {
		SendError(ctx, http.StatusInternalServerError, err)
		return
	}

	SendSuccess(ctx, http.StatusCreated, "Opening created successfully", createdId)
}

func ListOpeningsHandler(ctx *gin.Context) {
	openings, err := services.ListOpenings()
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err)
		return
	}

	parsedOpenings := handlerSchema.ParseListOpeningResponse(openings)
	SendSuccess(ctx, http.StatusOK, "Openings retrieved successfully", parsedOpenings)
}
