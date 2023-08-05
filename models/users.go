package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null;index"`
}
