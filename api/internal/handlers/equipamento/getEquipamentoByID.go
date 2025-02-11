package handlers

import (
	"net/http"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getEquipamentoResponse representa a resposta para a busca de um equipamento
type getEquipamentoResponse struct {
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

// @Summary Buscar um equipamento pelo ID
// @Description Retorna os detalhes de um equipamento específico
// @Tags Equipamentos
// @Accept json
// @Produce json
// @Param id path string true "ID do equipamento"
// @Success 200 {object} getEquipamentoResponse "Detalhes do equipamento"
// @Failure 400 {object} models.ErrorResponse "ID inválido"
// @Failure 404 {object} models.ErrorResponse "Equipamento não encontrado"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /equipamentos/{id} [get]
func GetEquipamentoByID(store db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Capturar o ID da URL
		idParam := ctx.Param("id")

		// Converter para UUID
		id, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "ID inválido"})
			return
		}

		// Buscar o equipamento no banco de dados
		equipamento, err := store.GetEquipamentoByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Equipamento não encontrado"})
			return
		}

		// Converter `Localizacao` de `[]byte` para `string` (se necessário)
		localizacaoStr := ""
		if loc, ok := equipamento.Localizacao.([]byte); ok {
			localizacaoStr = string(loc) // Convertendo de []byte para WKT
		}

		// Criar resposta formatada
		response := getEquipamentoResponse{
			ID:                    equipamento.ID,
			Nome:                  equipamento.Nome,
			Descricao:             models.NullString{String: equipamento.Descricao.String, Valid: equipamento.Descricao.Valid},
			Tipo:                  models.NullString{String: equipamento.Tipo.String, Valid: equipamento.Tipo.Valid},
			NumeroSerie:           equipamento.NumeroSerie.String,
			Modelo:                models.NullString{String: equipamento.Modelo.String, Valid: equipamento.Modelo.Valid},
			Fabricante:            models.NullString{String: equipamento.Fabricante.String, Valid: equipamento.Fabricante.Valid},
			Frequencia:            models.NullFloat64{Float64: equipamento.Frequencia.Float64, Valid: equipamento.Frequencia.Valid},
			DataCalibracao:        models.NullTime{String: equipamento.DataCalibracao.Time.Format(time.RFC3339), Valid: equipamento.DataCalibracao.Valid},
			DataUltimaManutencao:  models.NullTime{String: equipamento.DataUltimaManutencao.Time.Format(time.RFC3339), Valid: equipamento.DataUltimaManutencao.Valid},
			ResponsavelManutencao: models.NullString{String: equipamento.ResponsavelManutencao.String, Valid: equipamento.ResponsavelManutencao.Valid},
			DataFabricacao:        models.NullTime{String: equipamento.DataFabricacao.Time.Format(time.RFC3339), Valid: equipamento.DataFabricacao.Valid},
			DataAquisicao:         models.NullTime{String: equipamento.DataAquisicao.Time.Format(time.RFC3339), Valid: equipamento.DataAquisicao.Valid},
			TiposDados:            models.NullString{String: equipamento.TiposDados.String, Valid: equipamento.TiposDados.Valid},
			Notas:                 models.NullString{String: equipamento.Notas.String, Valid: equipamento.Notas.Valid},
			DataExpiracaoGarantia: models.NullTime{String: equipamento.DataExpiracaoGarantia.Time.Format(time.RFC3339), Valid: equipamento.DataExpiracaoGarantia.Valid},
			StatusOperacional:     equipamento.StatusOperacional.String,
			Localizacao:           localizacaoStr,
			Imagem:                models.NullString{String: equipamento.Imagem.String, Valid: equipamento.Imagem.Valid},
		}

		// Retornar resposta JSON
		ctx.JSON(http.StatusOK, response)
	}
}
