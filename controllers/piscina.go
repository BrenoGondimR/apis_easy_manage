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

func UpdatePiscinaByID(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da tratamento não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da tratamento inválido"})
		return
	}

	// Obtenha os dados da atualização da manutenção do corpo da solicitação.
	var updatedTratamento models.Tratamento
	if err := c.BindJSON(&updatedTratamento); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da tratamento"})
		return
	}

	// Utilize a função UpdateManutencaoByID da model para atualizar a manutenção no banco de dados.
	if err := models.UpdateTratamentoByID(utils.DB, id, &updatedTratamento); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar tratamento", "message": err.Error()})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "tratamento atualizada com sucesso"})
}

func GetTratamentoByIDAndPopulateForm(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da tratamento não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da tratamento inválido"})
		return
	}

	// Utilize a função GetManutencaoByID da model para buscar a manutenção pelo ID.
	tratamento, err := models.GetTratamentoByID(utils.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar tratamento", "message": err.Error()})
		return
	}

	// Preencha os valores do formulário com base nos dados da manutenção.
	form := &models.Tratamento{
		ID:             tratamento.ID,
		NomePiscineiro: tratamento.NomePiscineiro,
		NomeEmpresa:    tratamento.NomeEmpresa,
		Data:           tratamento.Data,
		Pha:            tratamento.Pha,
		Cloro:          tratamento.Cloro,
		Acidez:         tratamento.Acidez,
		Alcalinidade:   tratamento.Alcalinidade,
		Type:           tratamento.Type,
	}

	// Retorne os dados do formulário preenchidos como resposta JSON.
	c.JSON(200, gin.H{"data": form})
}

func GetAllTratamentos(c *gin.Context) {
	tratamentoModel := models.Tratamento{}

	tratamentos, err := tratamentoModel.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar tratamento"})
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
