package controller

import "github.com/gin-gonic/gin"

func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getTodos",
	})
}

func CreateTodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createTodo",
	})
}
