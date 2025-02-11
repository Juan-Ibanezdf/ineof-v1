package router

import (
	campanha "github.com/Juan-Ibanezdf/ineof-v1/internal/handlers/campanha"
	equipamento "github.com/Juan-Ibanezdf/ineof-v1/internal/handlers/equipamento"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initializeRoutes configura as rotas da API
func initializeRoutes(router *gin.Engine, r *Router) {
	v1 := router.Group("/api/v1")
	{
		// Rota Swagger UI
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		// Rota Campanhas
		v1.PUT("/campanhas/:id", campanha.UpdateCampanha(r.store))    // ✅ Adicionando o UPDATE
		v1.GET("/campanhas", campanha.ListCampanhas(r.store))         // ✅ Adicionando o GET All
		v1.GET("/campanhas/:id", campanha.GetCampanhaByID(r.store))   // ✅ Adicionando o GET por ID
		v1.POST("/campanhas", campanha.CreateCampanha(r.store))       // ✅ Passamos `store` para o handler corretamente
		v1.DELETE("/campanhas/:id", campanha.DeleteCampanha(r.store)) // ✅ Adicionando o DELETE

		// Rotas Equipamentos
		v1.POST("/equipamentos/:id", equipamento.CreateEquipamento(r.store))   // ✅ Adicionando o CREATE
		v1.PUT("/equipamentos/:id", equipamento.UpdateEquipamento(r.store))    // ✅ Adicionando o UPDATE
		v1.GET("/equipamentos", equipamento.ListEquipamentos(r.store))         // ✅ Adicionando o GET All
		v1.GET("/equipamentos/:id", equipamento.GetEquipamentoByID(r.store))   // ✅ Adicionando o GET por ID
		v1.DELETE("/equipamentos/:id", equipamento.DeleteEquipamento(r.store)) // ✅ Adicionando o DELETE
	}
}
