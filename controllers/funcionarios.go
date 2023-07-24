package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

// CreateFuncionario recebe as informações do front-end, cria um novo funcionário e o insere no banco de dados.
func CreateFuncionario(c *gin.Context) {
	funcionario := models.Funcionario{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct funcionario.
	if err := c.BindJSON(&funcionario); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar funcionário"})
		return
	}
	// Realize aqui qualquer validação ou tratamento adicional dos dados do funcionário, se necessário.

	// Utilize a função InsertFuncionario da model para inserir o funcionário no banco de dados.
	if err := funcionario.InsertFuncionario(utils.DB); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar funcionário"})
		return
	}
	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(201, gin.H{"message": "Funcionário criado com sucesso", "funcionario": funcionario})
}
