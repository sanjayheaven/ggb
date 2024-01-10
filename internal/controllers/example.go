package controllers

import (
	"go-gin-boilerplate/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var exampleService = new(services.ExampleService)

type ExampleController struct{}

// @Router /examples/createExample [post]
// @Description Create Example
// @Tags Example
// @Param data body string true "data"
// @Success 200 {object} object
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

// @Router /examples/getExample [get]
// @Description Get Example
// @Tags Example
// @Param exampleId query int true "the example id"
// @Success 200 {object} object
func (exampleController *ExampleController) GetExample(ctx *gin.Context) {

	exampleIdStr, ok := ctx.GetQuery("exampleId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro param"})
		return
	}

	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "erro param"})
		return
	}

	res := exampleService.GetExample(exampleId)

	if res == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}
