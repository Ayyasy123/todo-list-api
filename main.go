package main

import (
	"log"

	"github.com/Ayyasy123/todo-list-api/config"
	"github.com/Ayyasy123/todo-list-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetupUserRoutes(config.DB, r)
	routes.SetupChecklistRoutes(config.DB, r)
	routes.SetupItemRoutes(config.DB, r)

	log.Println("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
