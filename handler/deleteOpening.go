package handler

import (
	"fmt"
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/vi

// @Summary Delete Opening
// @Description Deleta uma vaga
// @Tags Vagas
// @Accept json
// @Produce json
// @Param id query string true "Identificacao da vaga"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequered("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("A vaga com o id: %s não foi encontrada.", id))
		return
	}

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Houve um erro ao tentar deletar a vaga com o id: %s", id))
		return
	}

	sendSuccess(ctx, "Vaga deletada", opening)
}
