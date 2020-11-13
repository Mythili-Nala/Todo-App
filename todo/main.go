package main

import (
	"todo/controllers"
	"todo/models"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	db := utils.Database()
	db.AutoMigrate(&models.Todo{})

	router := initRouter()
	router.Run()
}

func initRouter() *gin.Engine {
	router := gin.Default()

	controller := new(controllers.Controller)

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", controller.CreateTodo)
		v1.GET("/", controller.FetchAllTodo)
		v1.GET("/:id", controller.FetchSingleTodo)
		v1.PUT("/:id", controller.UpdateTodo)
		v1.DELETE("/:id", controller.DeleteTodo)
	}

	return router
}
