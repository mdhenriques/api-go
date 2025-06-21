package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdhenriques/api-go/database"
	"github.com/mdhenriques/api-go/models"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CreateUser godoc
// @Summary      Cria um novo usuário
// @Description  Endpoint para criar um novo usuário
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body CreateUserInput true "Dados do Usuário"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Input recebido:", input)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	user := models.User{
		Username:       input.Username,
		Email:          input.Email,
		HashedPassword: string(hashedPassword),
		Xp:             0,
		Coins:          0,
	}

	result := database.DB.Create(&user)
	fmt.Println("Resultado do Create:", result.Error, "RowsAffected:", result.RowsAffected)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario criado com sucesso",
		"user":    user,
	})
}
