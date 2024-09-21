package handler

import (
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("erro de validação: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("erro ao criar dados no banco de dados: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro ao criar tabela no banco")
		return
	}

	sendSuccess(ctx, "Vaga registrada!", opening)
}
