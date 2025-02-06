package tests

import (
	"context"
	"math/rand"
	"testing"
	"time"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Gera um número aleatório dentro do intervalo fornecido
func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Cria um dado aleatório do LIDAR WindCube para todas as altitudes
func createRandomLidarWindcubeDado(t *testing.T) db.LidarWindcubeDado {
	campanha := createRandomCampanha(t)
	equipamento := createRandomEquipamento(t)

	arg := db.CreateLidarWindcubeDadoParams{
		CampanhaID:    campanha.ID,
		EquipamentoID: equipamento.ID,
		Timestamp:     time.Now().UTC().Truncate(time.Microsecond),
		Posicao:       sqlNullString("Norte"),
		Temperatura:   sqlNullFloat64(randomFloat(15, 30)),
		WiperCount:    sqlNullInt32(int32(rand.Intn(100))),
	}

	altitudes := []int{40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 250, 300}

	for _, alt := range altitudes {
		setLidarData(&arg, alt, randomFloat(-20, 50), randomFloat(-10, 10), randomFloat(0, 5), randomFloat(2, 25), randomFloat(0, 360), randomFloat(-10, 10), randomFloat(-10, 10), randomFloat(-5, 5))
	}

	lidarDado, err := testQueries.CreateLidarWindcubeDado(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, lidarDado)

	require.Equal(t, arg.CampanhaID, lidarDado.CampanhaID)
	require.Equal(t, arg.EquipamentoID, lidarDado.EquipamentoID)
	require.Equal(t, arg.Timestamp, lidarDado.Timestamp.UTC().Truncate(time.Microsecond))
	require.Equal(t, arg.Posicao.String, lidarDado.Posicao.String)
	require.Equal(t, arg.Temperatura.Float64, lidarDado.Temperatura.Float64)
	require.Equal(t, arg.WiperCount.Int32, lidarDado.WiperCount.Int32)

	for _, alt := range altitudes {
		validateLidarData(t, lidarDado, alt)
	}

	return lidarDado
}

func setLidarData(arg *db.CreateLidarWindcubeDadoParams, alt int, cnr, radial, disp, speed, direction, xWind, yWind, zWind float64) {
	switch alt {
	case 40:
		arg.Cnr40m = sqlNullFloat64(cnr)
		arg.RadialWindSpeed40m = sqlNullFloat64(radial)
		arg.WindSpeedDisp40m = sqlNullFloat64(disp)
		arg.WindSpeed40m = sqlNullFloat64(speed)
		arg.WindDirection40m = sqlNullFloat64(direction)
		arg.XWind40m = sqlNullFloat64(xWind)
		arg.YWind40m = sqlNullFloat64(yWind)
		arg.ZWind40m = sqlNullFloat64(zWind)
	case 50:
		arg.Cnr50m = sqlNullFloat64(cnr)
		arg.RadialWindSpeed50m = sqlNullFloat64(radial)
		arg.WindSpeedDisp50m = sqlNullFloat64(disp)
		arg.WindSpeed50m = sqlNullFloat64(speed)
		arg.WindDirection50m = sqlNullFloat64(direction)
		arg.XWind50m = sqlNullFloat64(xWind)
		arg.YWind50m = sqlNullFloat64(yWind)
		arg.ZWind50m = sqlNullFloat64(zWind)
	case 300:
		arg.Cnr300m = sqlNullFloat64(cnr)
		arg.RadialWindSpeed300m = sqlNullFloat64(radial)
		arg.WindSpeedDisp300m = sqlNullFloat64(disp)
		arg.WindSpeed300m = sqlNullFloat64(speed)
		arg.WindDirection300m = sqlNullFloat64(direction)
		arg.XWind300m = sqlNullFloat64(xWind)
		arg.YWind300m = sqlNullFloat64(yWind)
		arg.ZWind300m = sqlNullFloat64(zWind)
	}
}

func validateLidarData(t *testing.T, dado db.LidarWindcubeDado, alt int) {
	switch alt {
	case 40:
		require.NotNil(t, dado.Cnr40m)
		require.NotNil(t, dado.RadialWindSpeed40m)
		require.NotNil(t, dado.WindSpeedDisp40m)
	case 50:
		require.NotNil(t, dado.Cnr50m)
		require.NotNil(t, dado.RadialWindSpeed50m)
		require.NotNil(t, dado.WindSpeedDisp50m)
	case 300:
		require.NotNil(t, dado.Cnr300m)
		require.NotNil(t, dado.RadialWindSpeed300m)
		require.NotNil(t, dado.WindSpeedDisp300m)
	}
}

func TestCreateLidarWindcubeDado(t *testing.T) {
	createRandomLidarWindcubeDado(t)
}

// ✅ Teste de busca por ID
func TestGetLidarWindcubeDadoByID(t *testing.T) {
	lidarDado1 := createRandomLidarWindcubeDado(t)
	lidarDado2, err := testQueries.GetLidarWindcubeDadoByID(context.Background(), lidarDado1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, lidarDado2)
	require.Equal(t, lidarDado1.ID, lidarDado2.ID)
	require.Equal(t, lidarDado1.CampanhaID, lidarDado2.CampanhaID)
	require.Equal(t, lidarDado1.EquipamentoID, lidarDado2.EquipamentoID)
	require.Equal(t, lidarDado1.Posicao.String, lidarDado2.Posicao.String)
	require.Equal(t, lidarDado1.Temperatura.Float64, lidarDado2.Temperatura.Float64)

}

// ✅ Teste de listagem de dados por campanha
func createRandomLidarWindcubeDadoWithCampanha(t *testing.T, campanhaID uuid.UUID) db.LidarWindcubeDado {
	equipamento := createRandomEquipamento(t)

	arg := db.CreateLidarWindcubeDadoParams{
		CampanhaID:    campanhaID, // Garantimos que usamos a mesma campanha
		EquipamentoID: equipamento.ID,
		Timestamp:     time.Now().UTC().Truncate(time.Microsecond),
		Posicao:       sqlNullString("Norte"),
		Temperatura:   sqlNullFloat64(randomFloat(15, 30)),
		WiperCount:    sqlNullInt32(int32(rand.Intn(100))),
	}

	altitudes := []int{40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 250, 300}

	for _, alt := range altitudes {
		setLidarData(&arg, alt, randomFloat(-20, 50), randomFloat(-10, 10), randomFloat(0, 5), randomFloat(2, 25), randomFloat(0, 360), randomFloat(-10, 10), randomFloat(-10, 10), randomFloat(-5, 5))
	}

	lidarDado, err := testQueries.CreateLidarWindcubeDado(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, lidarDado)

	return lidarDado
}

func TestListLidarWindcubeDadosByCampanha(t *testing.T) {
	// Criamos uma campanha fixa
	campanha := createRandomCampanha(t)

	// Criamos 3 registros de dados LIDAR vinculados à mesma campanha
	for i := 0; i < 3; i++ {
		createRandomLidarWindcubeDadoWithCampanha(t, campanha.ID)
	}

	// Buscamos os dados associados à campanha criada
	lidarDados, err := testQueries.ListLidarWindcubeDadosByCampanha(context.Background(), campanha.ID)

	require.NoError(t, err)
	require.NotEmpty(t, lidarDados, "Erro: Nenhum dado encontrado para a campanha %v", campanha.ID)

	t.Logf("Dados retornados: %+v", lidarDados) // Log dos dados retornados para depuração

	// Verificamos se os dados foram recuperados corretamente
	for _, dado := range lidarDados {
		require.NotEmpty(t, dado.ID)
		require.Equal(t, campanha.ID, dado.CampanhaID)
	}
}

// ✅ Teste de remoção de um dado do LIDAR
func TestDeleteLidarWindcubeDado(t *testing.T) {
	lidarDado := createRandomLidarWindcubeDado(t)
	err := testQueries.DeleteLidarWindcubeDado(context.Background(), lidarDado.ID)
	require.NoError(t, err)

	lidarDado2, err := testQueries.GetLidarWindcubeDadoByID(context.Background(), lidarDado.ID)
	require.Error(t, err)
	require.Empty(t, lidarDado2)
}
