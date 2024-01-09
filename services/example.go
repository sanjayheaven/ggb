package services

import (
	"go-gin-boilerplate/models"
)

type ExampleService struct{}

func (exampleService *ExampleService) CreateExample(data interface{}) *models.Example {

	example := models.Example{}
	models.DB.Create(&example)
	return &example

}
