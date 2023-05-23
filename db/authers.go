package db

import "gorm.io/gorm"

type Authors struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
