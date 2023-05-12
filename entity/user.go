package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"not null" binding:"required"`
	Email    string `gorm:"unique;not null" binding:"email,required"`
	Password string `gorm:"not null" binding:"required,min=6"`
	Role     string `gorm:"not null" binding:"required,oneof=admin member"`
}
