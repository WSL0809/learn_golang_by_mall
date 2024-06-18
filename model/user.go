package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

const (
	PasswordCost = 12
	Active = "active"
)

func (user *User) SetPassword (password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}