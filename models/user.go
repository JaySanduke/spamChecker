// package models

// type User struct {
// 	ID       uint      `gorm:"primaryKey"`
// 	Name     string    `gorm:"not null"`
// 	Phone    string    `gorm:"uniqueIndex;not null"`
// 	Email    *string   `gorm:"uniqueIndex"`
// 	Password string    `gorm:"not null"`
// 	Contacts []Contact `gorm:"foreignKey:UserID"`
// }

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Phone    string  `gorm:"uniqueIndex;not null"`
	Email    *string `gorm:"uniqueIndex"`
	Password string  `gorm:"not null"`
}
