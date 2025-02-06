package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// updateCampanhaResponse representa a resposta ao atualizar uma campanha
type updateCampanhaResponse struct {
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

// updateCampanhaRequest representa o payload para atualizar uma campanha existente
type updateCampanhaRequest struct {
	Nome        string            `json:"nome"`
	Imagem      models.NullString `json:"imagem"`
	DataInicio  models.NullTime   `json:"data_inicio"`
	DataFim     models.NullTime   `json:"data_fim"`
	Equipe      models.NullString `json:"equipe"`
	Localizacao string            `json:"localizacao"`
	Objetivos   models.NullString `json:"objetivos"`
	Contato     models.NullString `json:"contato"`
	Status      models.NullString `json:"status"`
	Notas       models.NullString `json:"notas"`
	Descricao   models.NullString `json:"descricao"`
}

// @Summary Atualizar uma campanha
// @Description Atualiza os dados de uma campanha existente pelo ID
// @Tags Campanhas
// @Accept json
// @Produce json
// @Param id path string true "ID da Campanha"
// @Param campanha body updateCampanhaRequest true "Dados da campanha a serem atualizados"
// @Success 200 {object} updateCampanhaResponse "Campanha atualizada com sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {object} models.ErrorResponse "Campanha não encontrada"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /campanhas/{id} [put]
func UpdateCampanha(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Buscar os dados atuais da campanha antes da atualização
		campanhaAtual, err := store.GetCampanhaByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Campanha não encontrada"})
			return
		}

		var req updateCampanhaRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		arg := db.UpdateCampanhaParams{ID: id}

		// Convertendo Localização corretamente
		var localizacaoAtual string
		if loc, ok := campanhaAtual.Localizacao.([]byte); ok {
			localizacaoAtual = string(loc) // Converte []byte para string
		} else {
			localizacaoAtual = "" // Se for nulo, mantém vazio
		}

		// Criar resposta com os dados antigos
		response := updateCampanhaResponse{
			ID:          id,
			Nome:        campanhaAtual.Nome,
			Imagem:      models.NullString{String: campanhaAtual.Imagem.String, Valid: campanhaAtual.Imagem.Valid},
			DataInicio:  models.NullTime{String: campanhaAtual.DataInicio.Time.Format(time.RFC3339), Valid: campanhaAtual.DataInicio.Valid},
			DataFim:     models.NullTime{String: campanhaAtual.DataFim.Time.Format(time.RFC3339), Valid: campanhaAtual.DataFim.Valid},
			Equipe:      models.NullString{String: campanhaAtual.Equipe.String, Valid: campanhaAtual.Equipe.Valid},
			Localizacao: localizacaoAtual, // Agora com conversão segura
			Objetivos:   models.NullString{String: campanhaAtual.Objetivos.String, Valid: campanhaAtual.Objetivos.Valid},
			Contato:     models.NullString{String: campanhaAtual.Contato.String, Valid: campanhaAtual.Contato.Valid},
			Status:      models.NullString{String: campanhaAtual.Status.String, Valid: campanhaAtual.Status.Valid},
			Notas:       models.NullString{String: campanhaAtual.Notas.String, Valid: campanhaAtual.Notas.Valid},
			Descricao:   models.NullString{String: campanhaAtual.Descricao.String, Valid: campanhaAtual.Descricao.Valid},
		}

		// Atualizar somente os campos enviados
		if req.Nome != "" {
			arg.Nome = req.Nome
			response.Nome = req.Nome
		}

		if req.Imagem.Valid {
			arg.Imagem = sql.NullString{String: req.Imagem.String, Valid: true}
			response.Imagem = req.Imagem
		}

		if req.DataInicio.Valid {
			parsedTime, err := time.Parse(time.RFC3339, req.DataInicio.String)
			if err == nil {
				arg.DataInicio = sql.NullTime{Time: parsedTime, Valid: true}
				response.DataInicio = models.NullTime{String: parsedTime.Format(time.RFC3339), Valid: true}
			}
		}

		if req.DataFim.Valid {
			parsedTime, err := time.Parse(time.RFC3339, req.DataFim.String)
			if err == nil {
				arg.DataFim = sql.NullTime{Time: parsedTime, Valid: true}
				response.DataFim = models.NullTime{String: parsedTime.Format(time.RFC3339), Valid: true}
			}
		}

		if req.Equipe.Valid {
			arg.Equipe = sql.NullString{String: req.Equipe.String, Valid: true}
			response.Equipe = req.Equipe
		}

		if req.Localizacao != "" {
			arg.StGeomfromtext = fmt.Sprintf("ST_GeomFromText('%s', 4326)", req.Localizacao)
			response.Localizacao = req.Localizacao
		}

		if req.Objetivos.Valid {
			arg.Objetivos = sql.NullString{String: req.Objetivos.String, Valid: true}
			response.Objetivos = req.Objetivos
		}

		if req.Contato.Valid {
			arg.Contato = sql.NullString{String: req.Contato.String, Valid: true}
			response.Contato = req.Contato
		}

		if req.Status.Valid {
			arg.Status = sql.NullString{String: req.Status.String, Valid: true}
			response.Status = req.Status
		}

		if req.Notas.Valid {
			arg.Notas = sql.NullString{String: req.Notas.String, Valid: true}
			response.Notas = req.Notas
		}

		if req.Descricao.Valid {
			arg.Descricao = sql.NullString{String: req.Descricao.String, Valid: true}
			response.Descricao = req.Descricao
		}

		// Executar atualização no banco
		_, err = store.UpdateCampanha(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Retornar a estrutura `updateCampanhaResponse`
		ctx.JSON(http.StatusOK, response)
	}
}
