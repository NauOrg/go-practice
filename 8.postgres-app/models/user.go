package models

import (
	"postgres-project/services"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Post  []Post
}

func init() {
	services.DB.AutoMigrate(&User{})
}
