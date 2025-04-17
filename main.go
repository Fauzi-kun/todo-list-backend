package main

import (
	"github.com/Fauzi-kun/todolist/config"
	"github.com/Fauzi-kun/todolist/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()	
	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	routes.TodoRoute(r)

	r.Run(":8080")
}