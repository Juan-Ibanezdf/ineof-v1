package router

import (
	db "github.com/Juan-Ibanezdf/ineof-v1/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Router estrutura para armazenar dependências
type Router struct {
	Engine *gin.Engine
	store  *db.Queries
}

// NewRouter cria uma nova instância do router
func NewRouter(store *db.SQLStore) *Router {
	router := gin.Default()
	router.Use(CORSConfig()) // Configurar CORS

	r := &Router{
		Engine: router,
		store:  store.Queries,
	}

	// Configurar as rotas e associar o router
	initializeRoutes(router, r)

	return r
}

// CORSConfig configura as permissões CORS
func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}
