package main

import (
	"gin-practice/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/todos", controller.GetTodos)
	r.POST("/todos", controller.CreateTodo)
	r.Run(":8080")
}
