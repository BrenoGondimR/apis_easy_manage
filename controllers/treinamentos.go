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

func GetTreinamentoByIDAndPopulateForm(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID do treinamento não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID do treinamento inválido"})
		return
	}

	// Utilize a função GetManutencaoByID da model para buscar a manutenção pelo ID.
	treinamento, err := models.GetTreinamentoByID(utils.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar treinamento", "message": err.Error()})
		return
	}

	// Preencha os valores do formulário com base nos dados da manutenção.
	form := &models.Treinamentos{
		ID:                        treinamento.ID,
		Treinamento:               treinamento.Treinamento,
		CargaHoraria:              treinamento.CargaHoraria,
		DataTreinamento:           treinamento.DataTreinamento,
		CaracteristicaTreinamento: treinamento.CaracteristicaTreinamento,
		Funcionarios:              treinamento.Funcionarios,
		Tipo:                      treinamento.Tipo,
		Observacoes:               treinamento.Observacoes,
		Status:                    treinamento.Status,
	}

	// Retorne os dados do formulário preenchidos como resposta JSON.
	c.JSON(200, gin.H{"data": form})
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

func UpdateTreinamentoByID(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da treinamento não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da treinamento inválido"})
		return
	}

	// Obtenha os dados da atualização da manutenção do corpo da solicitação.
	var updatedTreinamento models.Treinamentos
	if err := c.BindJSON(&updatedTreinamento); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da treinamento"})
		return
	}

	// Utilize a função UpdateManutencaoByID da model para atualizar a manutenção no banco de dados.
	if err := models.UpdateTreinamentoByID(utils.DB, id, &updatedTreinamento); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar treinamento", "message": err.Error()})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Treinamento atualizado com sucesso"})
}

func UpdateTreinamentosStatus(c *gin.Context) {
	treinamentos := models.Treinamentos{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct Manutencao.
	if err := c.BindJSON(&treinamentos); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações do Treinamento"})
		return
	}

	// Verifique se o ID da manutenção foi fornecido pelo front-end.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID do Treinamento não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID do Treinamento inválido"})
		return
	}

	// Utilize a função UpdateManutencao da model para atualizar o estado da manutenção no banco de dados.
	if err := models.UpdateTreinamentoStatus(utils.DB, id, treinamentos.Status); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar estado do Treinamento"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Status do Treinamento atualizado com sucesso"})
}
