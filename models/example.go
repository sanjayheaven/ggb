package models

import (
	"gorm.io/gorm"
)

type Example struct {
	gorm.Model

	Name   string // Name
	Status string // Status, active or inactive
}
