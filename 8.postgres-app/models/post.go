package models

import (
	"postgres-project/services"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}

func init() {
	services.DB.AutoMigrate(&Post{})
}
