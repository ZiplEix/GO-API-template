package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"text;not null;default:null"`
	Description string `json:"description" gorm:"text;default:null"`
	Completed   bool   `json:"completed" gorm:"not null;default:false"`
	UserID      uint   `json:"user_id" gorm:"not null;default:null"`
}
