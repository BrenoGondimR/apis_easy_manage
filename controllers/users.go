package controllers

import (
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"example/web-service-gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	// Inicialize uma instância da estrutura Usuarios com os dados do JSON recebido do frontend.
	usuario := models.Usuarios{}
	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	// Fazer o hash da senha antes de armazená-la no banco de dados.
	hashedPassword, err := helpers.HashPassword(usuario.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	usuario.Password = hashedPassword

	// Utilize a função CreateUser do modelo para inserir o usuário no banco de dados.
	createdID, err := usuario.CreateUser(utils.DB)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar usuário", "message": err.Error()})
		return
	}

	// O usuário foi criado com sucesso. Você pode retornar uma resposta de sucesso ao frontend.
	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso", "usuario_id": createdID})
}
