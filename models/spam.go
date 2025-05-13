package models

import "gorm.io/gorm"

type SpamReport struct {
	gorm.Model
	ReporterID  uint   `gorm:"index"`
	PhoneNumber string `gorm:"index;not null"`
}
