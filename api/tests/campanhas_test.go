package tests

import (
	"context"
	"testing"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// createRandomCampanha cria uma campanha aleatória para testes
func createRandomCampanha(t *testing.T) db.Campanha {
	arg := db.CreateCampanhaParams{
		ID:             uuid.New(),
		Nome:           "Campanha Teste",
		Imagem:         sqlNullString("https://example.com/image.jpg"),
		DataInicio:     sqlNullTime(time.Now().UTC().Truncate(time.Microsecond)),                  // Normaliza para UTC
		DataFim:        sqlNullTime(time.Now().AddDate(0, 1, 0).UTC().Truncate(time.Microsecond)), // +1 mês
		Equipe:         sqlNullString("Equipe Alpha"),
		StGeomfromtext: "SRID=4326;POINT(-46.625290 -23.533773)", // Localização válida
		Objetivos:      sqlNullString("Teste de campanha"),
		Contato:        sqlNullString("contato@example.com"),
		Status:         sqlNullString("Em Andamento"),
		Notas:          sqlNullString("Nota de teste"),
		Descricao:      sqlNullString("Descrição da campanha"),
	}

	campanha, err := testQueries.CreateCampanha(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, campanha)

	// Ajuste para timestamps normalizados
	require.Equal(t, arg.DataInicio.Time, campanha.DataInicio.Time.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.DataFim.Time, campanha.DataFim.Time.UTC().Truncate(time.Microsecond))

	// Verificações de dados
	require.Equal(t, arg.Nome, campanha.Nome)
	require.Equal(t, arg.Imagem.String, campanha.Imagem.String)
	require.Equal(t, arg.Equipe.String, campanha.Equipe.String)
	require.Equal(t, arg.Objetivos.String, campanha.Objetivos.String)
	require.Equal(t, arg.Contato.String, campanha.Contato.String)
	require.Equal(t, arg.Status.String, campanha.Status.String)
	require.Equal(t, arg.Notas.String, campanha.Notas.String)
	require.Equal(t, arg.Descricao.String, campanha.Descricao.String)

	return campanha
}

// TestCreateCampanha verifica a criação de uma campanha
func TestCreateCampanha(t *testing.T) {
	createRandomCampanha(t)
}

// TestUpdateCampanha verifica a atualização de uma campanha existente
func TestUpdateCampanha(t *testing.T) {
	campanha1 := createRandomCampanha(t)

	arg := db.UpdateCampanhaParams{
		ID:             campanha1.ID,
		Nome:           "Campanha Atualizada",
		Imagem:         sqlNullString("https://example.com/new-image.jpg"),
		DataInicio:     sqlNullTime(time.Now().UTC()),                  // Convertendo para UTC
		DataFim:        sqlNullTime(time.Now().AddDate(0, 2, 0).UTC()), // +2 meses
		Equipe:         sqlNullString("Equipe Beta"),
		StGeomfromtext: "POINT(-46.635290 -23.543773)",
		Objetivos:      sqlNullString("Objetivo atualizado"),
		Contato:        sqlNullString("newcontato@example.com"),
		Status:         sqlNullString("Finalizada"),
		Notas:          sqlNullString("Notas atualizadas"),
		Descricao:      sqlNullString("Descrição atualizada"),
	}

	campanha2, err := testQueries.UpdateCampanha(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, campanha2)

	require.Equal(t, arg.Nome, campanha2.Nome)
	require.Equal(t, arg.Imagem.String, campanha2.Imagem.String)

	// Convertendo para UTC antes da comparação para evitar erro de fuso horário
	require.WithinDuration(t, arg.DataInicio.Time.UTC(), campanha2.DataInicio.Time.UTC(), time.Millisecond)
	require.WithinDuration(t, arg.DataFim.Time.UTC(), campanha2.DataFim.Time.UTC(), time.Millisecond)

	require.Equal(t, arg.Equipe.String, campanha2.Equipe.String)
	require.Equal(t, arg.Objetivos.String, campanha2.Objetivos.String)
	require.Equal(t, arg.Contato.String, campanha2.Contato.String)
	require.Equal(t, arg.Status.String, campanha2.Status.String)
	require.Equal(t, arg.Notas.String, campanha2.Notas.String)
	require.Equal(t, arg.Descricao.String, campanha2.Descricao.String)
}

// TestGetCampanhaByID verifica a busca de uma campanha pelo ID
func TestGetCampanhaByID(t *testing.T) {
	campanha1 := createRandomCampanha(t)
	campanha2, err := testQueries.GetCampanhaByID(context.Background(), campanha1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, campanha2)

	require.Equal(t, campanha1.ID, campanha2.ID)
	require.Equal(t, campanha1.Nome, campanha2.Nome)
	require.Equal(t, campanha1.Imagem.String, campanha2.Imagem.String)
	require.Equal(t, campanha1.DataInicio.Time, campanha2.DataInicio.Time)
	require.Equal(t, campanha1.DataFim.Time, campanha2.DataFim.Time)
	require.Equal(t, campanha1.Equipe.String, campanha2.Equipe.String)
	require.Equal(t, campanha1.Objetivos.String, campanha2.Objetivos.String)
	require.Equal(t, campanha1.Contato.String, campanha2.Contato.String)
	require.Equal(t, campanha1.Status.String, campanha2.Status.String)
	require.Equal(t, campanha1.Notas.String, campanha2.Notas.String)
	require.Equal(t, campanha1.Descricao.String, campanha2.Descricao.String)
}

// TestListCampanhas verifica a listagem de campanhas
func TestListCampanhas(t *testing.T) {
	// Criamos algumas campanhas para garantir que a listagem funciona
	for i := 0; i < 5; i++ {
		createRandomCampanha(t)
	}

	campanhas, err := testQueries.ListCampanhas(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, campanhas)

	// Verificamos se os dados estão sendo retornados corretamente
	for _, campanha := range campanhas {
		require.NotEmpty(t, campanha.ID)
		require.NotEmpty(t, campanha.Nome)
	}
}

// TestDeleteCampanha verifica a exclusão de uma campanha
func TestDeleteCampanha(t *testing.T) {
	campanha1 := createRandomCampanha(t)
	err := testQueries.DeleteCampanha(context.Background(), campanha1.ID)
	require.NoError(t, err)

	campanha2, err := testQueries.GetCampanhaByID(context.Background(), campanha1.ID)
	require.Error(t, err)
	require.Empty(t, campanha2)
}
