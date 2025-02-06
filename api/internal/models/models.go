package models

import (
	"database/sql"
	"encoding/json"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/google/uuid"
)

// ErrorResponse estrutura para erros no Swagger
type ErrorResponse struct {
	Error string `json:"error"`
}

// NullString é um wrapper para sql.NullString
type NullString struct {
	String string `json:"string,omitempty"`
	Valid  bool   `json:"valid"`
}

// MarshalJSON para Swagger exibir corretamente
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// ToSQLNull converte para sql.NullString
func (ns NullString) ToSQLNull() sql.NullString {
	return sql.NullString{String: ns.String, Valid: ns.Valid}
}

// UnmarshalJSON customizado para NullString
func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &ns.String); err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

// NullTime é um wrapper para sql.NullTime
type NullTime struct {
	String string `json:"string,omitempty"`
	Valid  bool   `json:"valid"`
}

// MarshalJSON para Swagger exibir corretamente
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.String)
}

// ToSQLNull converte para sql.NullTime
func (nt NullTime) ToSQLNull() sql.NullTime {
	parsedTime, err := time.Parse(time.RFC3339, nt.String)
	if err != nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: parsedTime, Valid: nt.Valid}
}

// MarshalJSON para Swagger exibir corretamente
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float64)
}

// ToSQLNull converte para sql.NullFloat64
func (nf NullFloat64) ToSQLNull() sql.NullFloat64 {
	return sql.NullFloat64{Float64: nf.Float64, Valid: nf.Valid}
}

// UnmarshalJSON customizado para NullFloat64
func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nf.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &nf.Float64); err != nil {
		return err
	}
	nf.Valid = true
	return nil
}

// CampanhaResponse representa uma resposta serializável para Swagger
type CampanhaResponse struct {
	ID          uuid.UUID `json:"id"`
	Nome        string    `json:"nome"`
	Imagem      *string   `json:"imagem,omitempty"`
	DataInicio  *string   `json:"data_inicio,omitempty"`
	DataFim     *string   `json:"data_fim,omitempty"`
	Equipe      *string   `json:"equipe,omitempty"`
	Localizacao string    `json:"localizacao"`
	Objetivos   *string   `json:"objetivos,omitempty"`
	Contato     *string   `json:"contato,omitempty"`
	Status      *string   `json:"status,omitempty"`
	Notas       *string   `json:"notas,omitempty"`
	Descricao   *string   `json:"descricao,omitempty"`
}

// ConvertCampanha converte `db.Campanha` para `models.CampanhaResponse`
func ConvertCampanha(campanha db.Campanha) CampanhaResponse {
	localizacaoStr := ""
	if loc, ok := campanha.Localizacao.(string); ok {
		localizacaoStr = loc
	}
	return CampanhaResponse{
		ID:          campanha.ID,
		Nome:        campanha.Nome,
		Imagem:      nullStringToPtr(campanha.Imagem),
		DataInicio:  nullTimeToPtr(campanha.DataInicio),
		DataFim:     nullTimeToPtr(campanha.DataFim),
		Equipe:      nullStringToPtr(campanha.Equipe),
		Localizacao: localizacaoStr,
		Objetivos:   nullStringToPtr(campanha.Objetivos),
		Contato:     nullStringToPtr(campanha.Contato),
		Status:      nullStringToPtr(campanha.Status),
		Notas:       nullStringToPtr(campanha.Notas),
		Descricao:   nullStringToPtr(campanha.Descricao),
	}
}

// Funções auxiliares para converter `sql.NullString`, `sql.NullTime` e `sql.NullFloat64`
func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullTimeToPtr(nt sql.NullTime) *string {
	if nt.Valid {
		t := nt.Time.Format(time.RFC3339)
		return &t
	}
	return nil
}

func nullFloat64ToPtr(nf sql.NullFloat64) *float64 {
	if nf.Valid {
		return &nf.Float64
	}
	return nil
}

// EquipamentoResponse representa uma resposta serializável para Swagger
type EquipamentoResponse struct {
	ID                    uuid.UUID `json:"id"`
	Nome                  string    `json:"nome"`
	Descricao             *string   `json:"descricao,omitempty"`
	Tipo                  *string   `json:"tipo,omitempty"`
	NumeroSerie           *string   `json:"numero_serie,omitempty"`
	Modelo                *string   `json:"modelo,omitempty"`
	Fabricante            *string   `json:"fabricante,omitempty"`
	Frequencia            *float64  `json:"frequencia,omitempty"`
	DataCalibracao        *string   `json:"data_calibracao,omitempty"`
	DataUltimaManutencao  *string   `json:"data_ultima_manutencao,omitempty"`
	ResponsavelManutencao *string   `json:"responsavel_manutencao,omitempty"`
	DataFabricacao        *string   `json:"data_fabricacao,omitempty"`
	DataAquisicao         *string   `json:"data_aquisicao,omitempty"`
	TiposDados            *string   `json:"tipos_dados,omitempty"`
	Notas                 *string   `json:"notas,omitempty"`
	DataExpiracaoGarantia *string   `json:"data_expiracao_garantia,omitempty"`
	StatusOperacional     *string   `json:"status_operacional,omitempty"`
	Localizacao           string    `json:"localizacao"`
	Imagem                *string   `json:"imagem,omitempty"`
}

// ConvertEquipamento converte `db.Equipamento` para `models.EquipamentoResponse`
func ConvertEquipamento(equipamento db.Equipamento) EquipamentoResponse {
	localizacaoStr := ""
	if loc, ok := equipamento.Localizacao.(string); ok {
		localizacaoStr = loc
	}

	return EquipamentoResponse{
		ID:                    equipamento.ID,
		Nome:                  equipamento.Nome,
		Descricao:             nullStringToPtr(equipamento.Descricao),
		Tipo:                  nullStringToPtr(equipamento.Tipo),
		NumeroSerie:           nullStringToPtr(equipamento.NumeroSerie),
		Modelo:                nullStringToPtr(equipamento.Modelo),
		Fabricante:            nullStringToPtr(equipamento.Fabricante),
		Frequencia:            nullFloat64ToPtr(equipamento.Frequencia),
		DataCalibracao:        nullTimeToPtr(equipamento.DataCalibracao),
		DataUltimaManutencao:  nullTimeToPtr(equipamento.DataUltimaManutencao),
		ResponsavelManutencao: nullStringToPtr(equipamento.ResponsavelManutencao),
		DataFabricacao:        nullTimeToPtr(equipamento.DataFabricacao),
		DataAquisicao:         nullTimeToPtr(equipamento.DataAquisicao),
		TiposDados:            nullStringToPtr(equipamento.TiposDados),
		Notas:                 nullStringToPtr(equipamento.Notas),
		DataExpiracaoGarantia: nullTimeToPtr(equipamento.DataExpiracaoGarantia),
		StatusOperacional:     nullStringToPtr(equipamento.StatusOperacional),
		Localizacao:           localizacaoStr,
		Imagem:                nullStringToPtr(equipamento.Imagem),
	}
}

// NullFloat64 é um wrapper para sql.NullFloat64
type NullFloat64 struct {
	Float64 float64 `json:"float64,omitempty"`
	Valid   bool    `json:"valid"`
}
