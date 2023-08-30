package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFornecedor(c *gin.Context) {
	fornecedor := models.Fornecedor{}
	// BindJSON é utilizado para decodificar o JSON recebido do front-end para a struct fornecedor.
	if err := c.BindJSON(&fornecedor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar fornecedor"})
		return
	}

	// Verifique se todos os campos obrigatórios estão preenchidos
	if fornecedor.Fornecedor == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'fornecedor' não preenchido"})
		return
	}
	if fornecedor.Cidade == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'cidade' não preenchido"})
		return
	}
	if fornecedor.Contato == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'contato' não preenchido"})
		return
	}

	// Utilize a função InsertFornecedor da model para inserir o fornecedor no banco de dados.
	if err := fornecedor.InsertFornecedor(utils.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar fornecedor"})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(http.StatusCreated, gin.H{"message": "Fornecedor criado com sucesso", "fornecedor": fornecedor})
}

func GetAllFornecedores(c *gin.Context) {
	fornecedorModel := models.Fornecedor{}

	fornecedores, err := fornecedorModel.FindAll(utils.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar funcionários"})
		return
	}

	// Retorna os funcionários encontrados para o front-end
	c.JSON(200, gin.H{"data": fornecedores})
}
