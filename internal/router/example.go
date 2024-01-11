package routes

import (
	"go-gin-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

func LoadExampleRoutes(r *gin.Engine) *gin.RouterGroup {
	exampleController := new(controllers.ExampleController)
	example := r.Group("/examples")
	{
		example.POST("/createExample", exampleController.CreateExample)
		example.GET("/getExample", exampleController.GetExample)
		example.POST("/updateExample", exampleController.UpdateExample)
		example.POST("/deleteExample", exampleController.DeleteExample)
	}
	return example
}
