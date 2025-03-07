package routes

import (
	"github.com/automatedtomato/restful-task-api/handlers"
	"github.com/gin-gonic/gin"
)

// configure all the routes for app
func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", handlers.GetTasks)
			tasks.GET("/:id", handlers.GetTask)
			tasks.POST("", handlers.CreateTask)
			tasks.PUT("/:id", handlers.UpdateTask)
			tasks.DELETE("/:id", handlers.DeleteTask)
		}
	}
}
