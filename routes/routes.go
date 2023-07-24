package routes

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	// Funcionarios
	router.POST("/funcionarios/create", controllers.CreateFuncionario)
}
