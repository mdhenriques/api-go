package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mdhenriques/api-go/database"
	"github.com/mdhenriques/api-go/models"
)

type CreateTaskInput struct {
	Titulo        string     `json:"titulo" binding:"required"`
	Descricao     string     `json:"descricao"`
	PrazoEntrega  *time.Time `json:"prazo_entrega" format:"date-time"`
	TempoEstimado *float64   `json:"tempo_estimado`
}

// @Summary      Criar nova tarefa
// @Description  Cria uma nova tarefa para o usuário autenticado
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        task body controllers.CreateTaskInput true "Dados da nova tarefa"
// @Success      201 {object} models.TaskResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario nao autenticado"})
		return
	}

	userId := userIdInterface.(int)

	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		UserID:        uint(userId),
		Titulo:        input.Titulo,
		Descricao:     input.Descricao,
		PrazoEntrega:  input.PrazoEntrega,
		TempoEstimado: input.TempoEstimado,
	}

	result := database.DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, task)
}
