package handlers

import (
	"net/http"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// deleteCampanhaResponse representa a resposta para a exclusão de uma campanha
type deleteCampanhaResponse struct {
	Message string `json:"message"`
}

// @Summary Deletar uma campanha
// @Description Remove uma campanha do sistema pelo ID
// @Tags Campanhas
// @Accept json
// @Produce json
// @Param id path string true "ID da Campanha"
// @Success 200 {object} deleteCampanhaResponse "Mensagem de sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {object} models.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /campanhas/{id} [delete]
func DeleteCampanha(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Capturar o ID da campanha na URL
		idParam := ctx.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Verificar se a campanha existe antes de deletar
		_, err = store.GetCampanhaByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Campanha não encontrada"})
			return
		}

		// Deletar a campanha no banco de dados
		err = store.DeleteCampanha(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Retornar mensagem de sucesso
		response := deleteCampanhaResponse{Message: "Campanha deletada com sucesso"}
		ctx.JSON(http.StatusOK, response)
	}
}
