package controllers

import (
	"go-gin-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var exampleService = new(services.ExampleService)

type ExampleController struct{}

// @Router /examples/createExample [post]
// @Description Create Example
// @Tags Example
// @Param data body string true "data"
// @Success 200 {object} object
// @Return 200 {object} object
func (exampleController *ExampleController) CreateExample(ctx *gin.Context) {

	var data map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := exampleService.CreateExample(data)
	if res == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}
