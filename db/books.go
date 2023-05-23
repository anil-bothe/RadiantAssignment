package db

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
