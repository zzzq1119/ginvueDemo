package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(11);not null,unique"`
	Password string `gorm:"size:255;not null"`
}

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	result := db.Where("phone = ?", phone).First(&user)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
