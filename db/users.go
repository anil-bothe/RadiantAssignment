package db

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password"`
}
