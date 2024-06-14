package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
	PasswordDigest string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(100);not null"`
	Status string `gorm:"not null"`
	Avatar string `gorm:"type:varchar(100);not null"`
	Money string `gorm:"type:varchar(100);not null"`
}