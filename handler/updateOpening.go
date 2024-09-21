package handler

import (
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("erro de validação")
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequered(
			"id",
			"queryParameter",
		).Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "A vaga não foi encontrada")
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("erro ao atualizar a vaga: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro ao atualizar a vaga")
		return
	}

	sendSuccess(ctx, "vaga atualizada", opening)
}
