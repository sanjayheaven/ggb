package services

import (
	"go-gin-boilerplate/internal/models"
)

type ExampleService struct{}

func (exampleService *ExampleService) CreateExample(data map[string]interface{}) *models.Example {

	example := models.Example{
		Name: data["name"].(string),
	}

	res := models.DB.Create(&example)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}

	return &example

}

func (exampleService *ExampleService) GetExample(exampleId int) *models.Example {

	example := models.Example{}
	res := models.DB.First(&example, exampleId).Select("id, name, status")

	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}
	return &example
}

func (exampleService *ExampleService) UpdateExample(data map[string]interface{}) bool {

	example := models.Example{}
	example.ID = uint(data["exampleId"].(int))

	res := models.DB.Model(&example).Updates(data)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}

func (exampleService *ExampleService) DeleteExample(exampleId int) bool {

	example := models.Example{}
	res := models.DB.Delete(&example, exampleId)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
