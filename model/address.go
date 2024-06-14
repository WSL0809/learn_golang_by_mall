package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	user_id uint `gorm:"not null" `
	name string `gorm:"type:varchar(100);not null"`
	phone  string `gorm:"type:varchar(100);not null"`
	address string `gorm:"type:varchar(100);not null"`
}