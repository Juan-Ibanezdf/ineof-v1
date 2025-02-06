package handlers

import (
	"net/http"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Listar todas as campanhas
// @Description Retorna todas as campanhas registradas no sistema
// @Tags Campanhas
// @Accept json
// @Produce json
// @Success 200 {array} models.CampanhaResponse "Lista de campanhas"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /campanhas [get]
func ListCampanhas(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Buscar campanhas do banco de dados
		campanhasDB, err := store.ListCampanhas(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Converter db.Campanha para models.CampanhaResponse
		var campanhas []models.CampanhaResponse
		for _, campanha := range campanhasDB {
			campanhas = append(campanhas, models.ConvertCampanha(campanha))
		}

		// Retornar resposta JSON
		ctx.JSON(http.StatusOK, campanhas)
	}
}
