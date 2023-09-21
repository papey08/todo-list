package httpserver

import (
	"github.com/gin-gonic/gin"
	"todo-list/internal/app"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo-list/docs"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/task", addTask(a))
	r.GET("/task/:id", getTaskById(a))
	r.GET("/task", getTaskByText(a))
	r.PUT("/task/:id", updateTask(a))
	r.DELETE("/task/:id", deleteTask(a))
	r.GET("/task/by_status", getTasksByStatus(a))
	r.GET("/task/by_date", getTasksByDateAndStatus(a))
}
