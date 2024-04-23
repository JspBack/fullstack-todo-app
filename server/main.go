package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// todo represents a task in the todo list
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// GeoLocation represents geolocation data for an IP address
type GeoLocation struct {
	IP          string  `json:"query"`
	Country     string  `json:"country"`
	Region      string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Organization string  `json:"org"`
}

// todos stores the list of todos
var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a RESTful API", Completed: false},
	{ID: "3", Item: "Build a React App", Completed: false},
}

// getAllTodos returns all todos
func getAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// addTodos adds a new todo
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

// deleteTodos deletes a todo by ID
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

// getTodo retrieves a todo by ID
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

// toggleTodoStatus toggles the completed status of a todo
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

// getTodoByID retrieves a todo by ID
func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

// logIPAndGeoLocationToFile logs IP addresses and their geolocations to a file
func logIPAndGeoLocationToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		method := c.Request.Method
		if ip == "127.0.0.1" || ip == "::1" {
			// Log only IP address without geolocation for local IPs
			logLine := fmt.Sprintf("[%s] %s %s \n", time.Now().Format(time.RFC3339), ip, method)
			writeToLogFile(logLine)
			c.Next()
			return
		}

		geo, err := getGeoLocation(ip)
		if err != nil {
			fmt.Println("Error getting geolocation:", err)
		}
		logLine := fmt.Sprintf("[%s] %s - Geolocation: %+v\n", time.Now().Format(time.RFC3339), ip, geo)
		writeToLogFile(logLine)
		c.Next()
	}
}

// writeToLogFile writes log lines to a file
func writeToLogFile(logLine string) {
	f, err := os.OpenFile("ip_geo_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(logLine); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}

// getGeoLocation fetches geolocation data for an IP address
func getGeoLocation(ip string) (*GeoLocation, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geo GeoLocation
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, err
	}
	return &geo, nil
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(logIPAndGeoLocationToFile())

	router.GET("/todos", getAllTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.DELETE("/todos/:id", deleteTodos)
	router.Run("localhost:9090")
}
