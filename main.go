package main

import (
	"log"

	"github.com/automatedtomato/restful-task-api/routes"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/automatedtomato/restful-task-api/docs"
)

// @title 			Task API
// @version 		1.0
// @description 	A simple task management API
// @host 			localhost:8080
// @BasePath 		/api/v1

func main() {
	r := gin.Default()

	// CORS configuration
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// set up route group
	routes.SetupRoutes(r)

	// add Swagger doc endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// launch server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start over server: %v", err) // Fatalf: if error system shuts down itself
	}
}
