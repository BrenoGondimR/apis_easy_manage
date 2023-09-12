package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
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
