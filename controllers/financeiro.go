package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func CreateFinanceiro(c *gin.Context) {
	// Inicialize uma instância da estrutura Financeiro com os dados do JSON recebido do frontend.
	financeiro := models.Financeiro{}
	if err := c.BindJSON(&financeiro); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler dados financeiros"})
		return
	}

	// Verifique se pelo menos um dos campos 'ganhos' ou 'custos' está preenchido.
	if financeiro.Ganhos == 0 && financeiro.Custos == 0 {
		c.JSON(400, gin.H{"error": "Pelo menos um dos campos 'ganhos' ou 'custos' deve ser preenchido"})
		return
	}

	// Chame a função CreateFinanceiro para inserir os dados financeiros no banco de dados.
	createdID, err := financeiro.CreateFinanceiro(utils.DB)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar dados financeiros", "message": err.Error()})
		return
	}

	// Os dados financeiros foram criados com sucesso. Você pode retornar uma resposta de sucesso ao frontend.
	c.JSON(201, gin.H{"message": "Dados financeiros criados com sucesso", "financeiro_id": createdID})
}

func GetFinanceiroByIDAndPopulateForm(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da financeiro não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da financeiro inválido"})
		return
	}

	// Utilize a função GetManutencaoByID da model para buscar a manutenção pelo ID.
	financeiro, err := models.GetFinanceiroByID(utils.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar financeiro", "message": err.Error()})
		return
	}

	// Preencha os valores do formulário com base nos dados da financeiro.
	form := &models.Financeiro{
		ID:            financeiro.ID,
		Origem:        financeiro.Origem,
		Ganhos:        financeiro.Ganhos,
		TipoTransacao: financeiro.TipoTransacao,
		Custos:        financeiro.Custos,
		Status:        financeiro.Status,
		Data:          financeiro.Data,
	}

	// Retorne os dados do formulário preenchidos como resposta JSON.
	c.JSON(200, gin.H{"data": form})
}

func UpdateFinanceiroByID(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID da financeiro não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID da financeiro inválido"})
		return
	}

	// Obtenha os dados da atualização da manutenção do corpo da solicitação.
	var updatedFinanceiro models.Financeiro
	if err := c.BindJSON(&updatedFinanceiro); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da financeiro"})
		return
	}

	// Utilize a função UpdateManutencaoByID da model para atualizar a manutenção no banco de dados.
	if err := models.UpdateFinanceiroByID(utils.DB, id, &updatedFinanceiro); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar financeiro", "message": err.Error()})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "financeiro atualizado com sucesso"})
}

func GetTotaisMensais(c *gin.Context) {
	// Chame a função CalcularTotaisMensais para obter os totais mensais.
	totais, err := models.CalcularTotaisMensais(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao calcular totais mensais", "message": err.Error()})
		return
	}

	// Retorne apenas os valores dos totais mensais como resposta JSON.
	c.JSON(http.StatusOK, gin.H{
		"custos": totais.Custos,
		"ganhos": totais.Ganhos,
		"renda":  totais.Renda,
	})
}

func GetAllFinanceiro(c *gin.Context) {
	// Crie uma instância de *Manutencoes (ponteiro para Manutencoes)
	financeiro := &models.Financeiro{}

	// Invoque a função FindAll para buscar todas as manutenções no banco de dados.
	financeiroList, err := financeiro.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar Financeiro", "message": err.Error()})
		return
	}

	// Retorne a lista de manutenções encontradas como resposta JSON.
	c.JSON(200, gin.H{"data": financeiroList})
}

func GetAllCustos(c *gin.Context) {
	// Chame a função GetAllCustos para buscar todos os custos e calcular a soma.
	totalCustos, err := models.GetAllCustos(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar custos", "message": err.Error()})
		return
	}

	// Retorne o total de custos como resposta.
	c.JSON(http.StatusOK, gin.H{"totalCustos": totalCustos})
}

func GetAllGanhos(c *gin.Context) {
	// Chame a função GetAllCustos para buscar todos os custos e calcular a soma.
	totalGanhos, err := models.GetAllGanhos(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar ganhos", "message": err.Error()})
		return
	}

	// Retorne o total de custos como resposta.
	c.JSON(http.StatusOK, gin.H{"totalGanhos": totalGanhos})
}

func GetAllRenda(c *gin.Context) {
	// Chame a função GetAllRenda para calcular a renda (ganhos - custos).
	totalRenda, err := models.GetAllRenda(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao calcular a renda", "message": err.Error()})
		return
	}

	// Retorne o total de renda como resposta.
	c.JSON(http.StatusOK, gin.H{"totalRenda": totalRenda})
}
