package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// Funcionarios
	router.POST("/funcionarios/create", controllers.CreateFuncionario)
	router.GET("/funcionarios", controllers.GetAllFuncionarios)

	// Piscina
	router.POST("/piscina/create", controllers.CreateTratamento)
	router.POST("/fornecedores/create", controllers.CreateFornecedor)
	router.GET("/piscina", controllers.GetAllTratamentos)
	router.GET("/fornecedores", controllers.GetAllFornecedores)
	router.POST("/piscina/manutencao/create", controllers.CreateManutencao)
	router.GET("/piscina/manutencao", controllers.GetAllManutencoes)
	router.PUT("/piscina/estado/update/:manutID", controllers.UpdateManutencaoControllers)
}
