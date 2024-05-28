package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        string `json:"id" validate:"required,min=1"`
	Title     string `json:"title" validate:"required,min=1"`
	Completed bool   `json:"completed" validate:"required,min=1"`
}

var Todos = []Todo{}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173", "http://google.com"}
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	// Health Check endpoint
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "API is up and running")
	})

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, Todos)
	})

	router.PUT("/todos", func(c *gin.Context) {
		var tmpTodos []Todo

		if err := c.BindJSON(&tmpTodos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		Todos = tmpTodos

		c.Status(http.StatusNoContent)
	})

	return router
}

func main() {
	r := setupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
