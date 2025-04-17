package config

import (
	"fmt"
	"log"

	"github.com/Fauzi-kun/todolist/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := "host=localhost user=postgres password=fauzi dbname=todolist port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database:",err)
	}
	fmt.Println("Database connected.")
	DB.AutoMigrate(&models.Todo{})
}