package core

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string
	Name     string
	Email    string
	Password string
}
