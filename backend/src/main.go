package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	libUUID "github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ToDo struct {
	Id        string `json:"id" validate:"required,min=1"`
	Title     string `json:"title" validate:"required,min=1"`
	Completed bool   `json:"completed" validate:"required,min=1"`
}

var conn *redis.Client
var ctx = context.Background()

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
			toDos := []ToDo{}

			val, err := conn.Get(ctx, uuid).Result()
			if err == nil {
				if err = json.Unmarshal([]byte(val), &toDos); err != nil {
					log.Panic(err)
				}
			} else if err != redis.Nil {
				log.Panic(err)
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

			tmpToDosInByte, err := json.Marshal(tmpToDos)
			if err != nil {
				log.Panic(err)
			}

			uuid := c.GetString("uuid")
			if err = conn.Set(ctx, uuid, tmpToDosInByte, 7*24*time.Hour).Err(); err != nil {
				log.Panic(err)
			}

			c.Status(http.StatusNoContent)
		})
	}

	return router
}

func main() {
	// DB connection
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	conn = redis.NewClient(opt)

	defer func() {
		if err := conn.Close(); err != nil {
			log.Panic(err)
		}
	}()

	r := setupRouter()

	port := "8000"

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
