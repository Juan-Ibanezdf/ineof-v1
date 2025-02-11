package handlers

import (
	"net/http"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// deleteEquipamentoResponse representa a resposta para a exclusão de um equipamento
type deleteEquipamentoResponse struct {
	Message string `json:"message"`
}

// @Summary Deletar um equipamento
// @Description Remove um equipamento do sistema pelo ID
// @Tags Equipamentos
// @Accept json
// @Produce json
// @Param id path string true "ID do Equipamento"
// @Success 200 {object} deleteEquipamentoResponse "Mensagem de sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {object} models.ErrorResponse "Equipamento não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /equipamentos/{id} [delete]
func DeleteEquipamento(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Capturar o ID do equipamento na URL
		idParam := ctx.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Verificar se o equipamento existe antes de deletar
		_, err = store.GetEquipamentoByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Equipamento não encontrado"})
			return
		}

		// Deletar o equipamento no banco de dados
		err = store.DeleteEquipamento(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Retornar mensagem de sucesso
		response := deleteEquipamentoResponse{Message: "Equipamento deletado com sucesso"}
		ctx.JSON(http.StatusOK, response)
	}
}
