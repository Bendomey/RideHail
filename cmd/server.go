package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	// _ "github.com/heroku/x/hmetrics/onload"
)

const defaultPort = "8080"

func main() {
	//set port
	port := os.Getenv("PORT")
	log.Print(port)
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
