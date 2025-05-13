package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID uint   `gorm:"index"`
	Name   string `gorm:"not null"`
	Phone  string `gorm:"index;not null"`
}
