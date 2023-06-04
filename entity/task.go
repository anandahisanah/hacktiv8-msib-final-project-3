package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" binding:"required"`
	Description string `gorm:"not null" binding:"required"`
	CategoryID  uint
	Status      bool
	UserID      uint
}
