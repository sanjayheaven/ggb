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
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	res := exampleService.CreateExample(data)
	if res == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error"})
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
	exampleIdStr := ctx.Query("exampleId")
	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}
	res := exampleService.GetExample(exampleId)
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Router /examples/updateExample [post]
// @Description Update Example
// @Tags Example
// @Param data body string true "data"
// @Success 200 {object} object
func (exampleController *ExampleController) UpdateExample(ctx *gin.Context) {

	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	res := exampleService.UpdateExample(data)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success"})

}

// @Router /examples/deleteExample [post]
// @Description Delete Example
// @Tags Example
// @Param data body string true "data"
// @Success 200 {object} object
func (exampleController *ExampleController) DeleteExample(ctx *gin.Context) {

	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	exampleIdStr := data["exampleId"].(string)
	if exampleIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	exampleId, err := strconv.Atoi(exampleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	res := exampleService.DeleteExample(exampleId)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success"})

}
