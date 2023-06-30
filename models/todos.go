package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title" gorm:"text;not null;default:null"`
	Completed bool   `json:"completed" gorm:"not null;default:false"`
}
