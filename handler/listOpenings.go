package handler

import (
	"net/http"

	"github.com/MigueelLago/goportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Houve um erro ao tentar listar as vagas")
		return
	}

	sendSuccess(ctx, "Vagas listadas", openings)
}
