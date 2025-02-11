package handlers

import (
	"net/http"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Listar todos os equipamentos
// @Description Retorna todos os equipamentos registrados no sistema
// @Tags Equipamentos
// @Accept json
// @Produce json
// @Success 200 {array} models.EquipamentoResponse "Lista de equipamentos"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /equipamentos [get]
func ListEquipamentos(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Buscar equipamentos do banco de dados
		equipamentosDB, err := store.ListEquipamentos(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Converter db.Equipamento para models.EquipamentoResponse
		var equipamentos []models.EquipamentoResponse
		for _, equipamento := range equipamentosDB {
			equipamentos = append(equipamentos, models.ConvertEquipamento(equipamento))
		}

		// Retornar resposta JSON
		ctx.JSON(http.StatusOK, equipamentos)
	}
}
