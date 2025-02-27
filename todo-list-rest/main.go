package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var todos []Todo

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func addTodo(c *gin.Context) {
	var newTodo Todo
	err := c.BindJSON(&newTodo)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, todos)
}

func deleteTodo(c *gin.Context) {

	id := c.Param("id")

	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func doneTodo(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos[i].Done = true
			c.JSON(http.StatusOK, gin.H{"message": "Todo DONE successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", doneTodo)

	router.Run(":8080")
}
