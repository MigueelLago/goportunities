package handler

import (
	"fmt"
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequered("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"Não existe vaga com o id: %s",
			id,
		))
		return
	}

	sendSuccess(ctx, "Vaga encontrada", opening)
}
