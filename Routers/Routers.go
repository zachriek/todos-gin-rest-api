package Routers

import (
	"belajar-golang/latihan-2/gorm-gin-1/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/todos", Controllers.GetAllTodos)
		v1.POST("/todos", Controllers.AddNewTodo)
		v1.GET("/todos/:id", Controllers.GetTodo)
		v1.PUT("/todos/:id", Controllers.UpdateTodo)
		v1.DELETE("/todos/:id", Controllers.DeleteTodo)
	}

	return router
}
