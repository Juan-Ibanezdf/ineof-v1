package tests

import (
	"context"
	"testing"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Função auxiliar para gerar um número de série aleatório
func randomNumeroSerie() string {
	return "SN-" + uuid.New().String()[0:8] // Gera um UUID e pega os primeiros 8 caracteres
}

// Cria um equipamento aleatório para testes
func createRandomEquipamento(t *testing.T) db.Equipamento {
	arg := db.CreateEquipamentoParams{
		ID:                    uuid.New(),
		Nome:                  "Equipamento Teste",
		Descricao:             sqlNullString("Descrição do equipamento"),
		Tipo:                  sqlNullString("LIDAR"),
		NumeroSerie:           sqlNullString(randomNumeroSerie()), // Agora o número de série é único
		Modelo:                sqlNullString("Modelo X"),
		Fabricante:            sqlNullString("Fabricante Y"),
		Frequencia:            sqlNullFloat64(60.0),
		DataCalibracao:        sqlNullTime(time.Now().UTC().Truncate(time.Microsecond)),
		DataUltimaManutencao:  sqlNullTime(time.Now().AddDate(0, -3, 0).UTC().Truncate(time.Microsecond)), // -3 meses
		ResponsavelManutencao: sqlNullString("Técnico Z"),
		DataFabricacao:        sqlNullTime(time.Now().AddDate(-2, 0, 0).UTC().Truncate(time.Microsecond)), // -2 anos
		DataAquisicao:         sqlNullTime(time.Now().AddDate(-1, 0, 0).UTC().Truncate(time.Microsecond)), // -1 ano
		TiposDados:            sqlNullString("Velocidade do Vento, Direção do Vento"),
		Notas:                 sqlNullString("Equipamento em testes"),
		DataExpiracaoGarantia: sqlNullTime(time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Microsecond)), // +1 ano
		StatusOperacional:     sqlNullString("Operacional"),
		StGeomfromtext:        "SRID=4326;POINT(-46.635290 -23.543773)", // Localização geográfica
		Imagem:                sqlNullString("https://example.com/equipamento.jpg"),
	}

	equipamento, err := testQueries.CreateEquipamento(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, equipamento)

	// Normalizar timestamps para evitar falhas de precisão
	require.Equal(t, arg.DataCalibracao.Time, equipamento.DataCalibracao.Time.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.DataUltimaManutencao.Time, equipamento.DataUltimaManutencao.Time.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.DataFabricacao.Time, equipamento.DataFabricacao.Time.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.DataAquisicao.Time, equipamento.DataAquisicao.Time.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.DataExpiracaoGarantia.Time, equipamento.DataExpiracaoGarantia.Time.UTC().Truncate(time.Microsecond))

	// Verificações padrão
	require.Equal(t, arg.Nome, equipamento.Nome)
	require.Equal(t, arg.Descricao.String, equipamento.Descricao.String)
	require.Equal(t, arg.Tipo.String, equipamento.Tipo.String)
	require.Equal(t, arg.NumeroSerie.String, equipamento.NumeroSerie.String)
	require.Equal(t, arg.Modelo.String, equipamento.Modelo.String)
	require.Equal(t, arg.Fabricante.String, equipamento.Fabricante.String)
	require.Equal(t, arg.Frequencia.Float64, equipamento.Frequencia.Float64)
	require.Equal(t, arg.ResponsavelManutencao.String, equipamento.ResponsavelManutencao.String)
	require.Equal(t, arg.TiposDados.String, equipamento.TiposDados.String)
	require.Equal(t, arg.Notas.String, equipamento.Notas.String)
	require.Equal(t, arg.StatusOperacional.String, equipamento.StatusOperacional.String)
	require.Equal(t, arg.Imagem.String, equipamento.Imagem.String)

	return equipamento
}

// ✅ Testa a criação de um equipamento
func TestCreateEquipamento(t *testing.T) {
	createRandomEquipamento(t)
}

// ✅ Testa a atualização de um equipamento
func TestUpdateEquipamento(t *testing.T) {
	equipamento1 := createRandomEquipamento(t)

	arg := db.UpdateEquipamentoParams{
		ID:                    equipamento1.ID,
		Nome:                  "Equipamento Atualizado",
		Descricao:             sqlNullString("Nova descrição"),
		Tipo:                  sqlNullString("SODAR"),
		NumeroSerie:           equipamento1.NumeroSerie, // Mantém o mesmo Número de Série
		Modelo:                sqlNullString("Novo Modelo"),
		Fabricante:            sqlNullString("Novo Fabricante"),
		Frequencia:            sqlNullFloat64(75.0),
		DataCalibracao:        sqlNullTime(time.Now().UTC()),
		DataUltimaManutencao:  sqlNullTime(time.Now().AddDate(0, -6, 0).UTC()), // -6 meses
		ResponsavelManutencao: sqlNullString("Novo Técnico"),
		DataFabricacao:        sqlNullTime(time.Now().AddDate(-3, 0, 0).UTC()), // -3 anos
		DataAquisicao:         sqlNullTime(time.Now().AddDate(-2, 0, 0).UTC()), // -2 anos
		TiposDados:            sqlNullString("Pressão, Temperatura"),
		Notas:                 sqlNullString("Notas atualizadas"),
		DataExpiracaoGarantia: sqlNullTime(time.Now().AddDate(2, 0, 0).UTC()), // +2 anos
		StatusOperacional:     sqlNullString("Manutenção"),
		StGeomfromtext:        "SRID=4326;POINT(-46.650000 -23.550000)",
		Imagem:                sqlNullString("https://example.com/novo_equipamento.jpg"),
	}

	equipamento2, err := testQueries.UpdateEquipamento(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, equipamento2)

	require.Equal(t, arg.Nome, equipamento2.Nome)
	require.Equal(t, arg.Descricao.String, equipamento2.Descricao.String)
	require.Equal(t, arg.Tipo.String, equipamento2.Tipo.String)
	require.Equal(t, arg.NumeroSerie.String, equipamento2.NumeroSerie.String) // Agora não dá erro
	require.Equal(t, arg.Modelo.String, equipamento2.Modelo.String)
	require.Equal(t, arg.Fabricante.String, equipamento2.Fabricante.String)
	require.Equal(t, arg.StatusOperacional.String, equipamento2.StatusOperacional.String)
}

// ✅ Testa busca de um equipamento pelo ID
func TestGetEquipamentoByID(t *testing.T) {
	equipamento1 := createRandomEquipamento(t)
	equipamento2, err := testQueries.GetEquipamentoByID(context.Background(), equipamento1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, equipamento2)

	require.Equal(t, equipamento1.ID, equipamento2.ID)
	require.Equal(t, equipamento1.Nome, equipamento2.Nome)
	require.Equal(t, equipamento1.Descricao.String, equipamento2.Descricao.String)
	require.Equal(t, equipamento1.Tipo.String, equipamento2.Tipo.String)
	require.Equal(t, equipamento1.NumeroSerie.String, equipamento2.NumeroSerie.String)
	require.Equal(t, equipamento1.Modelo.String, equipamento2.Modelo.String)
	require.Equal(t, equipamento1.Fabricante.String, equipamento2.Fabricante.String)
	require.Equal(t, equipamento1.StatusOperacional.String, equipamento2.StatusOperacional.String)
}

// ✅ Testa listagem de equipamentos
func TestListEquipamentos(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomEquipamento(t)
	}

	equipamentos, err := testQueries.ListEquipamentos(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, equipamentos)

	for _, equipamento := range equipamentos {
		require.NotEmpty(t, equipamento.ID)
		require.NotEmpty(t, equipamento.Nome)
	}
}

// ✅ Testa remoção de um equipamento
func TestDeleteEquipamento(t *testing.T) {
	equipamento1 := createRandomEquipamento(t)
	err := testQueries.DeleteEquipamento(context.Background(), equipamento1.ID)
	require.NoError(t, err)

	equipamento2, err := testQueries.GetEquipamentoByID(context.Background(), equipamento1.ID)
	require.Error(t, err)
	require.Empty(t, equipamento2)
}
