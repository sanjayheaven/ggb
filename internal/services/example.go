package services

import "go-gin-boilerplate/internal/models"

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
