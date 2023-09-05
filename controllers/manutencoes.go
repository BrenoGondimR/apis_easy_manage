package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

func CreateManutencao(c *gin.Context) {
	// Inicialize uma instância da estrutura Manutencoes com os dados do JSON recebido do frontend.
	tratamento := models.Manutencoes{}
	if err := c.BindJSON(&tratamento); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler tratamento"})
		return
	}

	// Chame a função CreateManutencao para inserir a manutenção no banco de dados.
	createdID, err := tratamento.CreateManutencao(utils.DB)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar tratamento", "message": err.Error()})
		return
	}

	// A manutenção foi criada com sucesso. Você pode retornar uma resposta de sucesso ao frontend.
	c.JSON(201, gin.H{"message": "Tratamento criado com sucesso", "tratamento_id": createdID})
}

func GetAllManutencoes(c *gin.Context) {
	// Crie uma instância de *Manutencoes (ponteiro para Manutencoes)
	manutencoes := &models.Manutencoes{}

	// Invoque a função FindAll para buscar todas as manutenções no banco de dados.
	manutencoesList, err := manutencoes.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar manutenções", "message": err.Error()})
		return
	}

	// Retorne a lista de manutenções encontradas como resposta JSON.
	c.JSON(200, gin.H{"data": manutencoesList})
}
