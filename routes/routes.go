package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// Manutencoes
	router.POST("/manutencoes/create", controllers.CreateManutencao)
	router.GET("/manutencoes", controllers.GetAllManutencoes)

	// Piscina
	router.POST("/piscina/create", controllers.CreateTratamento)
	router.GET("/piscina", controllers.GetAllTratamentos)
	router.POST("/piscina/manutencao/create", controllers.CreateManutencaoPiscina)
	router.GET("/piscina/manutencao", controllers.GetAllManutencoesPiscina)
	router.PUT("/piscina/estado/update/:manutID", controllers.UpdateManutencaoControllers)

	// Fornecedores
	router.GET("/fornecedores", controllers.GetAllFornecedores)
	router.POST("/fornecedores/create", controllers.CreateFornecedor)

	// Treinamentos
	router.GET("/treinamentos", controllers.GetAllTreinamentos)
	router.POST("/treinamentos/create", controllers.CreateTreinamento)

	// Financeiro
	router.POST("/financeiro/create", controllers.CreateFinanceiro)
	router.GET("/financeiro/custos", controllers.GetAllCustos)
	router.GET("/financeiro/ganhos", controllers.GetAllGanhos)
	router.GET("/financeiro/renda", controllers.GetAllRenda)
}
