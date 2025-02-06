package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Carregar variáveis do .env
	err := godotenv.Load("../.env", "../../.env", "./.env")
	if err != nil {
		log.Println("⚠️ Aviso: Não foi possível carregar o .env para os testes, usando variáveis do sistema.")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("❌ Erro: Variáveis de ambiente do banco de dados não foram definidas corretamente.")
	}

	dbSource := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbPort, dbName)

	// Conectar ao banco de dados
	testDB, err = sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao banco de dados:", err)
	}

	testQueries = db.New(testDB)

	// Rodar os testes
	code := m.Run()

	// Fechar conexão
	testDB.Close()

	// Sair com código correto
	os.Exit(code)
}
