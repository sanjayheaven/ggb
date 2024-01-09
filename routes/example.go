package routes

import (
	"go-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Create example
// @Description Create example
// @Tags Example
// @Accept  json
// @Produce  json
// @Param   example body interface{}  true "example"
// @Success 200 {object} interface{}
// @Router /examples/createExample [post]
func LoadProductRoutes(r *gin.Engine) (*gin.RouterGroup, error) {
	exampleController := new(controllers.ExampleController)
	product := r.Group("/examples")
	{
		product.POST("/createExample", exampleController.CreateExample)
	}
	return product, nil
}
