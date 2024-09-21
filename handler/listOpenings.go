package handler

import (
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/vi

// @Summary List Opening
// @Description Lista todas as vagas
// @Tags Vagas
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Houve um erro ao tentar listar as vagas")
		return
	}

	sendSuccess(ctx, "Vagas listadas", openings)
}
