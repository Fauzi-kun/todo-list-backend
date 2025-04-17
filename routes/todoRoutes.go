package routes

import (
	"github.com/Fauzi-kun/todolist/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoute(router *gin.Engine){
	todos := router.Group("/todos")
	{
		todos.GET("/",controllers.GetTodos)
		todos.POST("/",controllers.CreateTodo)
		todos.PUT("/:id",controllers.UpdateTodo)
		todos.DELETE("/:id",controllers.DeleteTodo)
	}
}