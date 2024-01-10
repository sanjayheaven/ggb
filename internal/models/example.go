package models

type Example struct {
	BasicModel

	Name   string `json:"name"`                         // Name
	Status string `json:"status" gorm:"default:active"` // Status, active or inactive
}
