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
	router.POST("/fornecedores/create", controllers.CreateFornecedor)
	router.GET("/piscina", controllers.GetAllTratamentos)
	router.GET("/fornecedores", controllers.GetAllFornecedores)
	router.POST("/piscina/manutencao/create", controllers.CreateManutencaoPiscina)
	router.GET("/piscina/manutencao", controllers.GetAllManutencoesPiscina)
	router.PUT("/piscina/estado/update/:manutID", controllers.UpdateManutencaoControllers)
}
