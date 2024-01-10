package routes

import (
	"go-gin-boilerplate/internal/controllers"

	"github.com/gin-gonic/gin"
)

func LoadEXampleRoutes(r *gin.Engine) (*gin.RouterGroup, error) {
	exampleController := new(controllers.ExampleController)
	product := r.Group("/examples")
	{
		product.POST("/createExample", exampleController.CreateExample)
	}
	return product, nil
}
