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

// updateEquipamentoResponse representa a resposta ao atualizar um equipamento
type updateEquipamentoResponse struct {
	ID                    uuid.UUID          `json:"id"`
	Nome                  string             `json:"nome"`
	Descricao             models.NullString  `json:"descricao,omitempty"`
	Tipo                  models.NullString  `json:"tipo,omitempty"`
	NumeroSerie           string             `json:"numero_serie"`
	Modelo                models.NullString  `json:"modelo,omitempty"`
	Fabricante            models.NullString  `json:"fabricante,omitempty"`
	Frequencia            models.NullFloat64 `json:"frequencia,omitempty"`
	DataCalibracao        models.NullTime    `json:"data_calibracao,omitempty"`
	DataUltimaManutencao  models.NullTime    `json:"data_ultima_manutencao,omitempty"`
	ResponsavelManutencao models.NullString  `json:"responsavel_manutencao,omitempty"`
	DataFabricacao        models.NullTime    `json:"data_fabricacao,omitempty"`
	DataAquisicao         models.NullTime    `json:"data_aquisicao,omitempty"`
	TiposDados            models.NullString  `json:"tipos_dados,omitempty"`
	Notas                 models.NullString  `json:"notas,omitempty"`
	DataExpiracaoGarantia models.NullTime    `json:"data_expiracao_garantia,omitempty"`
	StatusOperacional     string             `json:"status_operacional"`
	Localizacao           string             `json:"localizacao,omitempty"`
	Imagem                models.NullString  `json:"imagem,omitempty"`
}

// @Summary Atualizar um equipamento
// @Description Atualiza os dados de um equipamento existente pelo ID
// @Tags Equipamentos
// @Accept json
// @Produce json
// @Param id path string true "ID do Equipamento"
// @Param equipamento body updateEquipamentoResponse true "Dados do equipamento a serem atualizados"
// @Success 200 {object} updateEquipamentoResponse "Equipamento atualizado com sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {object} models.ErrorResponse "Equipamento não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /equipamentos/{id} [put]
func UpdateEquipamento(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Buscar os dados atuais do equipamento antes da atualização
		equipamentoAtual, err := store.GetEquipamentoByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Equipamento não encontrado"})
			return
		}

		var req updateEquipamentoResponse
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		arg := db.UpdateEquipamentoParams{ID: id}

		// Verificar se a localização está em formato válido WKT (Well-Known Text)
		localizacaoStr := ""
		if req.Localizacao != "" {
			localizacaoStr = fmt.Sprintf("ST_GeomFromText('%s', 4326)", req.Localizacao)
		}
		// Criar resposta com os dados antigos
		response := updateEquipamentoResponse{
			ID:                    id,
			Nome:                  equipamentoAtual.Nome,
			Descricao:             models.NullString{String: equipamentoAtual.Descricao.String, Valid: equipamentoAtual.Descricao.Valid},
			Tipo:                  models.NullString{String: equipamentoAtual.Tipo.String, Valid: equipamentoAtual.Tipo.Valid},
			NumeroSerie:           equipamentoAtual.NumeroSerie.String,
			Modelo:                models.NullString{String: equipamentoAtual.Modelo.String, Valid: equipamentoAtual.Modelo.Valid},
			Fabricante:            models.NullString{String: equipamentoAtual.Fabricante.String, Valid: equipamentoAtual.Fabricante.Valid},
			Frequencia:            models.NullFloat64{Float64: equipamentoAtual.Frequencia.Float64, Valid: equipamentoAtual.Frequencia.Valid},
			DataCalibracao:        models.NullTime{String: equipamentoAtual.DataCalibracao.Time.Format(time.RFC3339), Valid: equipamentoAtual.DataCalibracao.Valid},
			DataUltimaManutencao:  models.NullTime{String: equipamentoAtual.DataUltimaManutencao.Time.Format(time.RFC3339), Valid: equipamentoAtual.DataUltimaManutencao.Valid},
			ResponsavelManutencao: models.NullString{String: equipamentoAtual.ResponsavelManutencao.String, Valid: equipamentoAtual.ResponsavelManutencao.Valid},
			DataFabricacao:        models.NullTime{String: equipamentoAtual.DataFabricacao.Time.Format(time.RFC3339), Valid: equipamentoAtual.DataFabricacao.Valid},
			DataAquisicao:         models.NullTime{String: equipamentoAtual.DataAquisicao.Time.Format(time.RFC3339), Valid: equipamentoAtual.DataAquisicao.Valid},
			TiposDados:            models.NullString{String: equipamentoAtual.TiposDados.String, Valid: equipamentoAtual.TiposDados.Valid},
			Notas:                 models.NullString{String: equipamentoAtual.Notas.String, Valid: equipamentoAtual.Notas.Valid},
			DataExpiracaoGarantia: models.NullTime{String: equipamentoAtual.DataExpiracaoGarantia.Time.Format(time.RFC3339), Valid: equipamentoAtual.DataExpiracaoGarantia.Valid},
			StatusOperacional:     equipamentoAtual.StatusOperacional.String,
			Localizacao:           localizacaoStr,
			Imagem:                models.NullString{String: equipamentoAtual.Imagem.String, Valid: equipamentoAtual.Imagem.Valid},
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

		if req.StatusOperacional != "" {
			arg.StatusOperacional = sql.NullString{String: req.StatusOperacional, Valid: true}
			response.StatusOperacional = req.StatusOperacional
		}

		if req.Localizacao != "" {
			arg.StGeomfromtext = fmt.Sprintf("ST_GeomFromText('%s', 4326)", req.Localizacao)
			response.Localizacao = req.Localizacao
		}

		// Executar atualização no banco
		_, err = store.UpdateEquipamento(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Retornar a estrutura `updateEquipamentoResponse`
		ctx.JSON(http.StatusOK, response)
	}
}
