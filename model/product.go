package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
	Category uint
	Title string `gorm:"type:varchar(100);not null"`
	Info string `gorm:"type:varchar(100);not null"`
	Price string `gorm:"type:varchar(100);not null"`
	DiscountPrice string `gorm:"type:varchar(100);not null"`
	OnSale bool `gorm:"default:false"`
	Num int
	BossId uint
	BossName string
	BossAvatar string
	ImgPath string
}