package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	libUUID "github.com/google/uuid"
)

type ToDo struct {
	Id        string `json:"id" validate:"required,min=1"`
	Title     string `json:"title" validate:"required,min=1"`
	Completed bool   `json:"completed" validate:"required,min=1"`
}

var ToDosList = map[string][]ToDo{}

// add the middleware function
func uuidMiddleware(c *gin.Context) {
	uuid := c.Param("uuid")

	err := libUUID.Validate(uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("uuid", uuid)
	c.Next()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost", "http://google.com"}
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	// Health Check endpoint
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "API is up and running")
	})

	authorized := router.Group("/todos/:uuid")
	authorized.Use(uuidMiddleware)
	{
		authorized.GET("", func(c *gin.Context) {
			uuid := c.GetString("uuid")
			toDos := ToDosList[uuid]
			if toDos == nil {
				toDos = []ToDo{}
			}
			c.JSON(http.StatusOK, toDos)
		})

		authorized.PUT("", func(c *gin.Context) {
			var tmpToDos []ToDo

			if err := c.BindJSON(&tmpToDos); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			uuid := c.GetString("uuid")
			ToDosList[uuid] = tmpToDos

			c.Status(http.StatusNoContent)
		})
	}

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
