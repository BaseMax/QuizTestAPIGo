package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null;index"`
	Description string `gorm:"type:TEXT;not null"`
}

type QuizResult struct {
	ID          uint      `gorm:"primaryKey;auto_increment"`
	Score       int       `gorm:"not null"`
	SubmittedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	QuizID      uint      `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	Quiz        Quiz      `gorm:"foreignKey:QuizID;references:ID;constraint:OnDelete:CASCADE"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type QuizResultDetail struct {
	ID         uint     `gorm:"primaryKey;auto_increment"`
	QuizID     uint     `gorm:"not null"`
	QuestionID uint     `gorm:"not null"`
	OptionID   uint     `gorm:"not null"`
	Quiz       Quiz     `gorm:"foreignKey:QuizID;references:ID;constraint:OnDelete:CASCADE"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID;constraint:OnDelete:CASCADE"`
	Option     Option   `gorm:"foreignKey:OptionID;references:ID;constraint:OnDelete:CASCADE"`
}
