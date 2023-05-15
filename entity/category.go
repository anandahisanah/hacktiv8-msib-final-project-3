package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type string `gorm:"not null" binding:"required"`
}
