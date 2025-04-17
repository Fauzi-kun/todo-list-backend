package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Fauzi-kun/todolist/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database:",err)
	}
	fmt.Println("Database connected.")
	DB.AutoMigrate(&models.Todo{})
}