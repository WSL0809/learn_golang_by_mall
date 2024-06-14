package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	username string `gorm:"type:varchar(100);not null"`
	password string `gorm:"type:varchar(100);not null"`
	avatar string `gorm:"type:varchar(100);not null"`
}