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

// createEquipamentoResponse representa a resposta ao criar um equipamento
type createEquipamentoResponse struct {
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

// createEquipamentoRequest representa o payload para criar um equipamento
type createEquipamentoRequest struct {
	Nome                  string             `json:"nome" binding:"required"`
	Descricao             models.NullString  `json:"descricao"`
	Tipo                  models.NullString  `json:"tipo"`
	NumeroSerie           string             `json:"numero_serie"`
	Modelo                models.NullString  `json:"modelo"`
	Fabricante            models.NullString  `json:"fabricante"`
	Frequencia            models.NullFloat64 `json:"frequencia"`
	DataCalibracao        models.NullTime    `json:"data_calibracao"`
	DataUltimaManutencao  models.NullTime    `json:"data_ultima_manutencao"`
	ResponsavelManutencao models.NullString  `json:"responsavel_manutencao"`
	DataFabricacao        models.NullTime    `json:"data_fabricacao"`
	DataAquisicao         models.NullTime    `json:"data_aquisicao"`
	TiposDados            models.NullString  `json:"tipos_dados"`
	Notas                 models.NullString  `json:"notas"`
	DataExpiracaoGarantia models.NullTime    `json:"data_expiracao_garantia"`
	StatusOperacional     string             `json:"status_operacional" binding:"required"`
	Localizacao           string             `json:"localizacao"`
	Imagem                models.NullString  `json:"imagem"`
}

// Função auxiliar para gerar um número de série aleatório
func randomNumeroSerie() string {
	return "SN-" + uuid.New().String()[0:8]
}

// @Summary Criar um novo equipamento
// @Description Cria um novo equipamento no sistema
// @Tags Equipamentos
// @Accept json
// @Produce json
// @Param equipamento body createEquipamentoRequest true "Dados do equipamento"
// @Success 200 {object} createEquipamentoResponse "Equipamento criado com sucesso"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /equipamentos [post]
func CreateEquipamento(server db.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createEquipamentoRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Se o número de série não for enviado, gera um automaticamente
		if req.NumeroSerie == "" {
			req.NumeroSerie = randomNumeroSerie()
		}
		// Criar `sql.NullString` corretamente para strings normais
		numeroSerieSQL := sql.NullString{String: req.NumeroSerie, Valid: req.NumeroSerie != ""}
		statusOperacionalSQL := sql.NullString{String: req.StatusOperacional, Valid: req.StatusOperacional != ""}

		// Preparar parâmetros para o banco
		arg := db.CreateEquipamentoParams{
			ID:                    uuid.New(),
			Nome:                  req.Nome,
			Descricao:             req.Descricao.ToSQLNull(),
			Tipo:                  req.Tipo.ToSQLNull(),
			NumeroSerie:           numeroSerieSQL,
			Modelo:                req.Modelo.ToSQLNull(),
			Fabricante:            req.Fabricante.ToSQLNull(),
			Frequencia:            req.Frequencia.ToSQLNull(),
			DataCalibracao:        req.DataCalibracao.ToSQLNull(),
			DataUltimaManutencao:  req.DataUltimaManutencao.ToSQLNull(),
			ResponsavelManutencao: req.ResponsavelManutencao.ToSQLNull(),
			DataFabricacao:        req.DataFabricacao.ToSQLNull(),
			DataAquisicao:         req.DataAquisicao.ToSQLNull(),
			TiposDados:            req.TiposDados.ToSQLNull(),
			Notas:                 req.Notas.ToSQLNull(),
			DataExpiracaoGarantia: req.DataExpiracaoGarantia.ToSQLNull(),
			StatusOperacional:     statusOperacionalSQL,
			StGeomfromtext:        req.Localizacao, // Mantém a estrutura para WKT
			Imagem:                req.Imagem.ToSQLNull(),
		}

		// Criar equipamento no banco
		equipamento, err := server.CreateEquipamento(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
			return
		}

		// Verificar se a localização está em formato válido WKT (Well-Known Text)
		localizacaoStr := ""
		if req.Localizacao != "" {
			localizacaoStr = fmt.Sprintf("ST_GeomFromText('%s', 4326)", req.Localizacao)
		}

		// Responder com os dados criados
		response := createEquipamentoResponse{
			ID:                    equipamento.ID,
			Nome:                  equipamento.Nome,
			Descricao:             models.NullString{String: equipamento.Descricao.String, Valid: equipamento.Descricao.Valid},
			Tipo:                  models.NullString{String: equipamento.Tipo.String, Valid: equipamento.Tipo.Valid},
			NumeroSerie:           equipamento.NumeroSerie.String, // ✅ Extraindo corretamente
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
			StatusOperacional:     equipamento.StatusOperacional.String, // ✅ Extraindo corretamente
			Localizacao:           localizacaoStr,
			Imagem:                models.NullString{String: equipamento.Imagem.String, Valid: equipamento.Imagem.Valid},
		}

		ctx.JSON(http.StatusOK, response)
	}
}
