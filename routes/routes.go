package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// Manutencoes
	router.POST("/manutencoes/create", controllers.CreateManutencao)
	router.GET("/manutencoes", controllers.GetAllManutencoes)
	router.PUT("/manutencoes/status/update/:manutID", controllers.UpdateManutencaoStatus)
	router.GET("/manutencoes/edit/get/:manutID", controllers.GetManutencaoByIDAndPopulateForm)
	router.PUT("/manutencoes/edit/update/:manutID", controllers.UpdateManutencaoByID)

	// Piscina
	router.POST("/piscina/create", controllers.CreateTratamento)
	router.GET("/piscina", controllers.GetAllTratamentos)
	router.POST("/piscina/manutencao/create", controllers.CreateManutencaoPiscina)
	router.GET("/piscina/manutencao", controllers.GetAllManutencoesPiscina)
	router.PUT("/piscina/estado/update/:manutID", controllers.UpdateManutencaoControllers)
	router.GET("/piscina/edit/get/:manutID", controllers.GetTratamentoByIDAndPopulateForm)
	router.PUT("/piscina/edit/update/:manutID", controllers.UpdatePiscinaByID)

	// Fornecedores
	router.GET("/fornecedores", controllers.GetAllFornecedores)
	router.POST("/fornecedores/create", controllers.CreateFornecedor)
	router.GET("/fornecedores/edit/get/:manutID", controllers.GetFornecedorByIDAndPopulateForm)
	router.PUT("/fornecedores/edit/update/:manutID", controllers.UpdateFornecedorByID)

	// Treinamentos
	router.GET("/treinamentos", controllers.GetAllTreinamentos)
	router.POST("/treinamentos/create", controllers.CreateTreinamento)
	router.PUT("/treinamentos/status/update/:manutID", controllers.UpdateTreinamentosStatus)
	router.GET("/treinamentos/edit/get/:manutID", controllers.GetTreinamentoByIDAndPopulateForm)
	router.PUT("/treinamentos/edit/update/:manutID", controllers.UpdateTreinamentoByID)

	// Financeiro
	router.POST("/financeiro/create", controllers.CreateFinanceiro)
	router.GET("/financeiro/custos", controllers.GetAllCustos)
	router.GET("/financeiro/ganhos", controllers.GetAllGanhos)
	router.GET("/financeiro/renda", controllers.GetAllRenda)
	router.GET("/financeiro/todos", controllers.GetAllFinanceiro)
	router.GET("/financeiro/cgr", controllers.GetTotaisMensais)
	router.GET("/financeiro/edit/get/:manutID", controllers.GetFinanceiroByIDAndPopulateForm)
	router.PUT("/financeiro/edit/update/:manutID", controllers.UpdateFinanceiroByID)

}
