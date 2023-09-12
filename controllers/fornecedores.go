package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateFornecedorByID(c *gin.Context) {
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
	var updatedFornecedor models.Fornecedor
	if err := c.BindJSON(&updatedFornecedor); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao ler informações da Manutenção"})
		return
	}

	// Utilize a função UpdateManutencaoByID da model para atualizar a manutenção no banco de dados.
	if err := models.UpdateFornecedorByID(utils.DB, id, &updatedFornecedor); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao atualizar Manutenção", "message": err.Error()})
		return
	}

	// Retorne uma resposta de sucesso ao front-end.
	c.JSON(200, gin.H{"message": "Fornecedor atualizada com sucesso"})
}

func GetFornecedorByIDAndPopulateForm(c *gin.Context) {
	// Obtenha o ID da manutenção da URL.
	idStr := c.Param("manutID")
	if idStr == "" {
		c.JSON(400, gin.H{"error": "ID do fornecedor não fornecido"})
		return
	}

	// Converta o ID da string para um tipo ObjectID.
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID do fornecedor inválido"})
		return
	}

	// Utilize a função GetManutencaoByID da model para buscar a manutenção pelo ID.
	fornecedores, err := models.GetFornecedorByID(utils.DB, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao buscar fornecedor", "message": err.Error()})
		return
	}

	// Preencha os valores do formulário com base nos dados da manutenção.
	form := &models.Fornecedor{
		ID:         fornecedores.ID,
		Fornecedor: fornecedores.Fornecedor,
		Cidade:     fornecedores.Cidade,
		Contato:    fornecedores.Contato,
		Nota:       fornecedores.Nota,
	}

	// Retorne os dados do formulário preenchidos como resposta JSON.
	c.JSON(200, gin.H{"data": form})
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
