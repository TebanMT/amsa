package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name           string
	LastName       string
	SecondLastName string
	Username       string
	Password       string
}
