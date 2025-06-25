package controllers

import (
	"net/http"
	"strconv"
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

// @Summary      Deletar tarefa
// @Description  Deleta uma tarefa do usuário autenticado com base no ID
// @Tags         tasks
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path int true "ID da Tarefa"
// @Success      204 "Tarefa deletada com sucesso"
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}
	userId := userIdInterface.(int)

	idParam := c.Param("id")
	taskId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar tarefa"})
		return
	}

	c.Status(http.StatusNoContent)
}
