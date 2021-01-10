package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// load environmwnt variables in here
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//set port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//setting up gin
	r := gin.Default()

	// r.Use(utils.GinContextToContextMiddleware())
	// r.POST("/", graphqlHandler(srv))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.GET("/graphql", playgroundHandler())
	r.Run()
}
