package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
)

// InitDB inicializa a conexão com o banco e retorna a store
func InitDB() *db.SQLStore {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Carregar credenciais do banco de dados
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	// Criar string de conexão com o PostgreSQL
	dbSource := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbPort, dbName)

	// Conectar ao banco de dados
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}

	// Criar a store do `sqlc`
	store := db.NewStore(conn)
	return store
}
