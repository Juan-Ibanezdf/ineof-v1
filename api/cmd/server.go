package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/Juan-Ibanezdf/ineof-v1/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Server representa o servidor da API
type Server struct {
	store  *db.SQLStore
	router *router.Router
}

// NewServer inicializa o servidor com a conexão do banco de dados
func NewServer() *Server {
	// Carregar variáveis de ambiente
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Obter configurações do banco
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	// Criar string de conexão
	dbSource := fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbPort, dbName)

	// Conectar ao banco de dados
	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}

	// Criar a store do `sqlc`
	store := db.NewStore(conn)

	// Criar o router com a store
	r := router.NewRouter(store)

	return &Server{
		store:  store,
		router: r,
	}
}

// Start inicia o servidor na porta especificada
func (s *Server) Start() error {
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080" // Porta padrão
	}

	log.Printf("Servidor rodando na porta %s", apiPort)
	return s.router.Engine.Run(":" + apiPort)
}
