package handler

import (
	"fmt"
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/vi

// @Summary Show Opening
// @Description Lista uma nova vaga via ID
// @Tags Vagas
// @Accept json
// @Produce json
// @Param id query string true "Identificacao da vaga"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
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
