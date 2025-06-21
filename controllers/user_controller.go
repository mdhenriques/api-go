package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/mdhenriques/api-go/database"
	"github.com/mdhenriques/api-go/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao carregar arquivo .env")
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

type LoginInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// @Summary      Login de usuário
// @Description  Faz login e retorna um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials body LoginInput true "Credenciais de Login"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario ou senha nao incorretos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario ou senha incorretos"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": user.ID,
	"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString((jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nao foi possivel gerar o token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}


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
