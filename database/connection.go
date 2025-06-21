package database

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mdhenriques/api-go/models"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

	fmt.Println("Conectando com DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	
	
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	

	DB = db
	fmt.Println("Banco conectado com sucesso!")




	err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Erro ao fazer AutoMigrate:", err)
    }
}