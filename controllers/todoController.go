package controllers

import (
	"net/http"

	"github.com/Fauzi-kun/todolist/config"
	"github.com/Fauzi-kun/todolist/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context){
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK,todos)
}
func CreateTodo(c *gin.Context){
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusOK,todo)
}
func UpdateTodo(c *gin.Context){
	id := c.Param("id")
	var todo models.Todo
	if err := config.DB.First(&todo,id).Error; err != nil{
		c.JSON(http.StatusNotFound,gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&todo)
	c.JSON(http.StatusOK,todo)
}
func DeleteTodo(c *gin.Context){
	id := c.Param("id")
	if err := config.DB.Delete(&models.Todo{},id).Error;err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to delete"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Todo deleted"})
}