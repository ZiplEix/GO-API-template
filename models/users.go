package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"text;not null;default:null; unique"`
	Password string `json:"password" gorm:"text;not null;default:null"`
}
