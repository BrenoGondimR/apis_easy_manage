package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateManutencaoStatus(c *gin.Context) {
	manutencao := models.Manutencoes{}
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
	if err := models.UpdateManutencaoStatus(utils.DB, id, manutencao.Status); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar estado da Manutenção"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Estado da Manutenção atualizado com sucesso"})
}

func UpdateManutencaoByID(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
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

	// Obtenha os dados da atualização da manutenção do corpo da solicitação.
	var updatedManutencao models.Manutencoes
	if err := c.BindJSON(&updatedManutencao); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da Manutenção"})
		return
	}

	// Utilize a função UpdateManutencaoByID da model para atualizar a manutenção no banco de dados.
	if err := models.UpdateManutencaoByID(utils.DB, id, &updatedManutencao); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar Manutenção", "message": err.Error()})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Manutenção atualizada com sucesso"})
}

func GetManutencaoByIDAndPopulateForm(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
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

	// Utilize a função GetManutencaoByID da model para buscar a manutenção pelo ID.
	manutencao, err := models.GetManutencaoByID(utils.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar Manutenção", "message": err.Error()})
		return
	}

	// Preencha os valores do formulário com base nos dados da manutenção.
	form := &models.Manutencoes{
		ID:             manutencao.ID,
		Local:          manutencao.Local,
		Tipo:           manutencao.Tipo,
		DataOcorrencia: manutencao.DataOcorrencia,
		Servico:        manutencao.Servico,
		Descricao:      manutencao.Descricao,
		Status:         manutencao.Status,
	}

	// Retorne os dados do formulário preenchidos como resposta JSON.
	c.JSON(200, gin.H{"data": form})
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
