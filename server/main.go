package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type todo struct {
	
		ID        string `json:"id"`
		Item      string `json:"item"`
		Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a RESTful API", Completed: false},
	{ID: "3", Item: "Build a React App", Completed: false},
}


func getAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo.ID = uuid.New().String()
	newTodo.Completed = false
	todos = append([]todo{newTodo}, todos...)
	context.IndentedJSON(http.StatusCreated, newTodo)
	fmt.Printf("New todo added: %+v\n", newTodo)
}


func deleteTodos(context *gin.Context) {
	id := context.Param("id")
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	t, err := getTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	t.Completed = !t.Completed
	c.IndentedJSON(http.StatusOK, t)
}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{ "PATCH", "GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin","Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        // AllowOriginFunc: func(origin string) bool {
        //     return origin == "https://github.com"
        // },
        MaxAge: 12 * time.Hour,
    }))
	router.GET("/todos", getAllTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.DELETE("/todos/:id", deleteTodos)
	router.Run("localhost:9090")
}


