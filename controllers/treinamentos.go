package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateTreinamento insere um novo treinamento na coleção de Treinamentos.
func CreateTreinamento(c *gin.Context) {
	treinamento := models.Treinamentos{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct Treinamentos.
	if err := c.ShouldBindJSON(&treinamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler treinamento"})
		return
	}
	// Utilize a função CreateTreinamento da model para inserir o treinamento no banco de dados.
	treinamentoID := treinamento.CreateTreinamento(utils.DB)
	if treinamentoID == primitive.NilObjectID {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar treinamento"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end com o ID do treinamento criado.
	c.JSON(http.StatusCreated, gin.H{"message": "Treinamento criado com sucesso", "treinamento_id": treinamentoID.Hex()})
}

func GetAllTreinamentos(c *gin.Context) {
	treinamentosModels := models.Treinamentos{}

	tratamentos, err := treinamentosModels.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar treinamentos"})
		return
	}

	// Retorna os funcionários encontrados para o front-end
	c.JSON(200, gin.H{"data": tratamentos})
}
