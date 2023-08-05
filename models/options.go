package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Text       string   `gorm:"type:TEXT;not null;index"`
	IsCorrect  bool     `gorm:"not null;default:false"`
	QuestionID uint     `gorm:"not null"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnDelete:CASCADE"`
}
