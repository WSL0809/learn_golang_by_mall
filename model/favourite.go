package model

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserId"`
	ProductId uint `gorm:"not null"`
	Product Product `gorm:"foreignKey:ProductId"`
	Boss User `gorm:"foreignKey:BossId"`
	BossId uint `gorm:"not null"`
}