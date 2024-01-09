package controllers

import (
	"go-gin-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var exampleService = new(services.ExampleService)

type ExampleController struct{}

func (exampleController *ExampleController) CreateExample(ctx *gin.Context) {

	data := ctx.Request.Body
	res := exampleService.CreateExample(data)
	ctx.JSON(http.StatusOK, gin.H{})

}
