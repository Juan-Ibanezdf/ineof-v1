package handlers

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// createCampanhaResponse representa a resposta ao criar uma campanha
type createCampanhaResponse struct {
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

// createCampanhaRequest representa o payload para criar uma campanha
type createCampanhaRequest struct {
	Nome        string            `json:"nome" binding:"required"`
	Imagem      models.NullString `json:"imagem"`
	DataInicio  models.NullTime   `json:"data_inicio"`
	DataFim     models.NullTime   `json:"data_fim"`
	Equipe      models.NullString `json:"equipe"`
	Localizacao string            `json:"localizacao" binding:"required"`
	Objetivos   models.NullString `json:"objetivos"`
	Contato     models.NullString `json:"contato"`
	Status      models.NullString `json:"status"`
	Notas       models.NullString `json:"notas"`
	Descricao   models.NullString `json:"descricao"`
}

// @Summary Criar uma nova campanha
// @Description Cria uma nova campanha no sistema
// @Tags Campanhas
// @Accept json
// @Produce json
// @Param campanha body createCampanhaRequest true "Dados da campanha"
// @Success 200 {object} createCampanhaResponse "Campanha criada com sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /campanhas [post]
func CreateCampanha(server db.Querier) gin.HandlerFunc { // ✅ Passamos `server` como argumento
	return func(ctx *gin.Context) {
		var req createCampanhaRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		arg := db.CreateCampanhaParams{
			ID:             uuid.New(),
			Nome:           req.Nome,
			Imagem:         req.Imagem.ToSQLNull(),
			DataInicio:     req.DataInicio.ToSQLNull(),
			DataFim:        req.DataFim.ToSQLNull(),
			Equipe:         req.Equipe.ToSQLNull(),
			StGeomfromtext: req.Localizacao,
			Objetivos:      req.Objetivos.ToSQLNull(),
			Contato:        req.Contato.ToSQLNull(),
			Status:         req.Status.ToSQLNull(),
			Notas:          req.Notas.ToSQLNull(),
			Descricao:      req.Descricao.ToSQLNull(),
		}

		// ✅ Usando `server.store` corretamente
		campanha, err := server.CreateCampanha(ctx, arg) // ✅ Correto
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Verificar se a localização está em formato válido WKT (Well-Known Text)
		localizacaoStr := ""
		if req.Localizacao != "" {
			localizacaoStr = fmt.Sprintf("ST_GeomFromText('%s', 4326)", req.Localizacao)
		}

		response := createCampanhaResponse{
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

		ctx.JSON(http.StatusOK, response)
	}
}
