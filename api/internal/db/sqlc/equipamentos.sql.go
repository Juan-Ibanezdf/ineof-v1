// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: equipamentos.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createEquipamento = `-- name: CreateEquipamento :one
INSERT INTO equipamentos (id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, ST_GeomFromText($18, 4326), $19)
RETURNING id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem
`

type CreateEquipamentoParams struct {
	ID                    uuid.UUID       `json:"id"`
	Nome                  string          `json:"nome"`
	Descricao             sql.NullString  `json:"descricao"`
	Tipo                  sql.NullString  `json:"tipo"`
	NumeroSerie           sql.NullString  `json:"numero_serie"`
	Modelo                sql.NullString  `json:"modelo"`
	Fabricante            sql.NullString  `json:"fabricante"`
	Frequencia            sql.NullFloat64 `json:"frequencia"`
	DataCalibracao        sql.NullTime    `json:"data_calibracao"`
	DataUltimaManutencao  sql.NullTime    `json:"data_ultima_manutencao"`
	ResponsavelManutencao sql.NullString  `json:"responsavel_manutencao"`
	DataFabricacao        sql.NullTime    `json:"data_fabricacao"`
	DataAquisicao         sql.NullTime    `json:"data_aquisicao"`
	TiposDados            sql.NullString  `json:"tipos_dados"`
	Notas                 sql.NullString  `json:"notas"`
	DataExpiracaoGarantia sql.NullTime    `json:"data_expiracao_garantia"`
	StatusOperacional     sql.NullString  `json:"status_operacional"`
	StGeomfromtext        interface{}     `json:"st_geomfromtext"`
	Imagem                sql.NullString  `json:"imagem"`
}

// Criar um novo equipamento
func (q *Queries) CreateEquipamento(ctx context.Context, arg CreateEquipamentoParams) (Equipamento, error) {
	row := q.db.QueryRowContext(ctx, createEquipamento,
		arg.ID,
		arg.Nome,
		arg.Descricao,
		arg.Tipo,
		arg.NumeroSerie,
		arg.Modelo,
		arg.Fabricante,
		arg.Frequencia,
		arg.DataCalibracao,
		arg.DataUltimaManutencao,
		arg.ResponsavelManutencao,
		arg.DataFabricacao,
		arg.DataAquisicao,
		arg.TiposDados,
		arg.Notas,
		arg.DataExpiracaoGarantia,
		arg.StatusOperacional,
		arg.StGeomfromtext,
		arg.Imagem,
	)
	var i Equipamento
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Tipo,
		&i.NumeroSerie,
		&i.Modelo,
		&i.Fabricante,
		&i.Frequencia,
		&i.DataCalibracao,
		&i.DataUltimaManutencao,
		&i.ResponsavelManutencao,
		&i.DataFabricacao,
		&i.DataAquisicao,
		&i.TiposDados,
		&i.Notas,
		&i.DataExpiracaoGarantia,
		&i.StatusOperacional,
		&i.Localizacao,
		&i.Imagem,
	)
	return i, err
}

const deleteEquipamento = `-- name: DeleteEquipamento :exec
DELETE FROM equipamentos WHERE id = $1
`

// Deletar um equipamento
func (q *Queries) DeleteEquipamento(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEquipamento, id)
	return err
}

const getEquipamentoByID = `-- name: GetEquipamentoByID :one
SELECT id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem FROM equipamentos WHERE id = $1
`

// Buscar um equipamento pelo ID
func (q *Queries) GetEquipamentoByID(ctx context.Context, id uuid.UUID) (Equipamento, error) {
	row := q.db.QueryRowContext(ctx, getEquipamentoByID, id)
	var i Equipamento
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Tipo,
		&i.NumeroSerie,
		&i.Modelo,
		&i.Fabricante,
		&i.Frequencia,
		&i.DataCalibracao,
		&i.DataUltimaManutencao,
		&i.ResponsavelManutencao,
		&i.DataFabricacao,
		&i.DataAquisicao,
		&i.TiposDados,
		&i.Notas,
		&i.DataExpiracaoGarantia,
		&i.StatusOperacional,
		&i.Localizacao,
		&i.Imagem,
	)
	return i, err
}

const listEquipamentos = `-- name: ListEquipamentos :many
SELECT id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem FROM equipamentos ORDER BY nome ASC
`

// Listar todos os equipamentos
func (q *Queries) ListEquipamentos(ctx context.Context) ([]Equipamento, error) {
	rows, err := q.db.QueryContext(ctx, listEquipamentos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Equipamento{}
	for rows.Next() {
		var i Equipamento
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Descricao,
			&i.Tipo,
			&i.NumeroSerie,
			&i.Modelo,
			&i.Fabricante,
			&i.Frequencia,
			&i.DataCalibracao,
			&i.DataUltimaManutencao,
			&i.ResponsavelManutencao,
			&i.DataFabricacao,
			&i.DataAquisicao,
			&i.TiposDados,
			&i.Notas,
			&i.DataExpiracaoGarantia,
			&i.StatusOperacional,
			&i.Localizacao,
			&i.Imagem,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEquipamento = `-- name: UpdateEquipamento :one
UPDATE equipamentos
SET 
    nome = $2,
    descricao = $3,
    tipo = $4,
    numero_serie = $5,
    modelo = $6,
    fabricante = $7,
    frequencia = $8,
    data_calibracao = $9,
    data_ultima_manutencao = $10,
    responsavel_manutencao = $11,
    data_fabricacao = $12,
    data_aquisicao = $13,
    tipos_dados = $14,
    notas = $15,
    data_expiracao_garantia = $16,
    status_operacional = $17,
    localizacao = ST_GeomFromText($18, 4326),
    imagem = $19
WHERE id = $1
RETURNING id, nome, descricao, tipo, numero_serie, modelo, fabricante, frequencia, data_calibracao, data_ultima_manutencao, responsavel_manutencao, data_fabricacao, data_aquisicao, tipos_dados, notas, data_expiracao_garantia, status_operacional, localizacao, imagem
`

type UpdateEquipamentoParams struct {
	ID                    uuid.UUID       `json:"id"`
	Nome                  string          `json:"nome"`
	Descricao             sql.NullString  `json:"descricao"`
	Tipo                  sql.NullString  `json:"tipo"`
	NumeroSerie           sql.NullString  `json:"numero_serie"`
	Modelo                sql.NullString  `json:"modelo"`
	Fabricante            sql.NullString  `json:"fabricante"`
	Frequencia            sql.NullFloat64 `json:"frequencia"`
	DataCalibracao        sql.NullTime    `json:"data_calibracao"`
	DataUltimaManutencao  sql.NullTime    `json:"data_ultima_manutencao"`
	ResponsavelManutencao sql.NullString  `json:"responsavel_manutencao"`
	DataFabricacao        sql.NullTime    `json:"data_fabricacao"`
	DataAquisicao         sql.NullTime    `json:"data_aquisicao"`
	TiposDados            sql.NullString  `json:"tipos_dados"`
	Notas                 sql.NullString  `json:"notas"`
	DataExpiracaoGarantia sql.NullTime    `json:"data_expiracao_garantia"`
	StatusOperacional     sql.NullString  `json:"status_operacional"`
	StGeomfromtext        interface{}     `json:"st_geomfromtext"`
	Imagem                sql.NullString  `json:"imagem"`
}

// Atualizar um equipamento existente
func (q *Queries) UpdateEquipamento(ctx context.Context, arg UpdateEquipamentoParams) (Equipamento, error) {
	row := q.db.QueryRowContext(ctx, updateEquipamento,
		arg.ID,
		arg.Nome,
		arg.Descricao,
		arg.Tipo,
		arg.NumeroSerie,
		arg.Modelo,
		arg.Fabricante,
		arg.Frequencia,
		arg.DataCalibracao,
		arg.DataUltimaManutencao,
		arg.ResponsavelManutencao,
		arg.DataFabricacao,
		arg.DataAquisicao,
		arg.TiposDados,
		arg.Notas,
		arg.DataExpiracaoGarantia,
		arg.StatusOperacional,
		arg.StGeomfromtext,
		arg.Imagem,
	)
	var i Equipamento
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Descricao,
		&i.Tipo,
		&i.NumeroSerie,
		&i.Modelo,
		&i.Fabricante,
		&i.Frequencia,
		&i.DataCalibracao,
		&i.DataUltimaManutencao,
		&i.ResponsavelManutencao,
		&i.DataFabricacao,
		&i.DataAquisicao,
		&i.TiposDados,
		&i.Notas,
		&i.DataExpiracaoGarantia,
		&i.StatusOperacional,
		&i.Localizacao,
		&i.Imagem,
	)
	return i, err
}
