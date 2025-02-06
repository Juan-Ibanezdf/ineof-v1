package handlers

import (
	"net/http"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getCampanhaResponse representa a resposta para a busca de uma campanha
type getCampanhaResponse struct {
	ID          uuid.UUID         `json:"id"`
	Nome        string            `json:"nome"`
	Imagem      models.NullString `json:"imagem,omitempty"`
	DataInicio  models.NullTime   `json:"data_inicio,omitempty"`
	DataFim     models.NullTime   `json:"data_fim,omitempty"`
	Equipe      models.NullString `json:"equipe,omitempty"`
	Localizacao string            `json:"localizacao"`
	Objetivos   models.NullString `json:"objetivos,omitempty"`
	Contato     models.NullString `json:"contato,omitempty"`
	Status      models.NullString `json:"status,omitempty"`
	Notas       models.NullString `json:"notas,omitempty"`
	Descricao   models.NullString `json:"descricao,omitempty"`
}

// @Summary Buscar uma campanha pelo ID
// @Description Retorna os detalhes de uma campanha específica
// @Tags Campanhas
// @Accept json
// @Produce json
// @Param id path string true "ID da campanha"
// @Success 200 {object} getCampanhaResponse "Detalhes da campanha"
// @Failure 400 {object} models.ErrorResponse "ID inválido"
// @Failure 404 {object} models.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /campanhas/{id} [get]
func GetCampanhaByID(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Pegar o ID da URL manualmente
		idParam := ctx.Param("id")

		// Converter para UUID
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Buscar a campanha no banco de dados
		campanha, err := store.GetCampanhaByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Campanha não encontrada"})
			return
		}
		localizacaoStr := ""
		if loc, ok := campanha.Localizacao.(string); ok {
			localizacaoStr = loc
		}
		// Criar resposta formatada
		response := getCampanhaResponse{
			ID:          campanha.ID,
			Nome:        campanha.Nome,
			Imagem:      models.NullString{String: campanha.Imagem.String, Valid: campanha.Imagem.Valid},
			DataInicio:  models.NullTime{String: campanha.DataInicio.Time.Format(time.RFC3339), Valid: campanha.DataInicio.Valid},
			DataFim:     models.NullTime{String: campanha.DataFim.Time.Format(time.RFC3339), Valid: campanha.DataFim.Valid},
			Equipe:      models.NullString{String: campanha.Equipe.String, Valid: campanha.Equipe.Valid},
			Localizacao: localizacaoStr,
			Objetivos:   models.NullString{String: campanha.Objetivos.String, Valid: campanha.Objetivos.Valid},
			Contato:     models.NullString{String: campanha.Contato.String, Valid: campanha.Contato.Valid},
			Status:      models.NullString{String: campanha.Status.String, Valid: campanha.Status.Valid},
			Notas:       models.NullString{String: campanha.Notas.String, Valid: campanha.Notas.Valid},
			Descricao:   models.NullString{String: campanha.Descricao.String, Valid: campanha.Descricao.Valid},
		}

		// Retornar resposta JSON
		ctx.JSON(http.StatusOK, response)
	}
}
