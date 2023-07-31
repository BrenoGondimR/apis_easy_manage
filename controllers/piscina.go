package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
)

func CreateTratamento(c *gin.Context) {
	tratamento := models.Tratamento{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct funcionario.
	if err := c.BindJSON(&tratamento); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler tratamento"})
		return
	}

	// Utilize a função InsertFuncionario da model para inserir o funcionário no banco de dados.
	if err := tratamento.InsertTratamento(utils.DB); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar tratamento"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(201, gin.H{"message": "Tratamento criado com sucesso", "tratamento": tratamento})
}

func CreateManutencao(c *gin.Context) {
	manutencao := models.Manutencao{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct funcionario.
	if err := c.BindJSON(&manutencao); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler Manutenção"})
		return
	}

	// Utilize a função InsertFuncionario da model para inserir o funcionário no banco de dados.
	if err := manutencao.InsertManutencao(utils.DB); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar Manutenção"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(201, gin.H{"message": "Manutenção criada com sucesso", "Manutenção": manutencao})
}

func GetAllTratamentos(c *gin.Context) {
	tratamentoModel := models.Tratamento{}

	tratamentos, err := tratamentoModel.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar funcionários"})
		return
	}

	// Retorna os funcionários encontrados para o front-end
	c.JSON(200, gin.H{"data": tratamentos})
}
