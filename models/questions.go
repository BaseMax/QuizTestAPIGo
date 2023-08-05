package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text   string `gorm:"type:TEXT;not null;index"`
	QuizID uint    `gorm:"not null"`
	Quiz   Quiz   `gorm:"foreignKey:QuizID;references:ID;constraint:OnDelete:CASCADE"`
}
