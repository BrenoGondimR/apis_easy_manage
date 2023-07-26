package controllers

import (
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

func CreateFuncionario(c *gin.Context) {
	funcionario := models.Funcionario{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct funcionario.
	if err := c.BindJSON(&funcionario); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar funcionário"})
		return
	}

	// Fazer o hash da senha antes de armazená-la no banco de dados
	hashSenha, err := helpers.HashPassword(funcionario.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar funcionário"})
		return
	}
	funcionario.Password = hashSenha

	// Utilize a função InsertFuncionario da model para inserir o funcionário no banco de dados.
	if err := funcionario.InsertFuncionario(utils.DB); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar funcionário"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(201, gin.H{"message": "Funcionário criado com sucesso", "funcionario": funcionario})
}

func GetAllFuncionarios(c *gin.Context) {
	funcionarioModel := models.Funcionario{}

	funcionarios, err := funcionarioModel.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar funcionários"})
		return
	}

	// Retorna os funcionários encontrados para o front-end
	c.JSON(200, gin.H{"data": funcionarios})
}
