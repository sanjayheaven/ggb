package routes

import (
	"go-gin-boilerplate/internal/controllers"
	"go-gin-boilerplate/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var exampleController = new(controllers.ExampleController)

func LoadExampleRoutes(r *gin.Engine) *gin.RouterGroup {

	example := r.Group("/examples")
	example.Use(middlewares.Jwt())

	{
		example.POST("/createExample", exampleController.CreateExample)
		example.GET("/getExample", exampleController.GetExample)
		example.POST("/updateExample", exampleController.UpdateExample)
		example.POST("/deleteExample", exampleController.DeleteExample)
	}
	return example
}
