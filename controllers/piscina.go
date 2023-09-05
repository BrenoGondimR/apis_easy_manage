package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateManutencaoPiscina(c *gin.Context) {
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

func UpdateManutencaoControllers(c *gin.Context) {
	manutencao := models.Manutencao{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct Manutencao.
	if err := c.BindJSON(&manutencao); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da Manutenção"})
		return
	}

	// Verifique se o ID da manutenção foi fornecido pelo front-end.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da Manutenção não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da Manutenção inválido"})
		return
	}

	// Utilize a função UpdateManutencao da model para atualizar o estado da manutenção no banco de dados.
	if err := models.UpdateManutencao(utils.DB, id, manutencao.Estado); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar estado da Manutenção"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Estado da Manutenção atualizado com sucesso"})
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

func GetAllManutencoesPiscina(c *gin.Context) {
	manutencaoModel := models.Manutencao{}

	manutencoes, err := manutencaoModel.FindAllManutencao(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar Manutenções"})
		return
	}

	// Retorna os funcionários encontrados para o front-end
	c.JSON(200, gin.H{"data": manutencoes})
}
