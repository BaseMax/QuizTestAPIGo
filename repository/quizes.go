package repository

import (
	"github.com/BaseMax/QuizTestAPIGo/models"
	"gorm.io/gorm"
)

type quizRepository struct {
	db *gorm.DB
}

type QuizRepository interface {
	Create(q *models.Quiz) error
	GetQuizByID(QuizID uint) (models.Quiz, error)
	GetAllQuizzes() ([]models.Quiz, error)
	Update(quiz *models.Quiz) error
	Delete(quiz *models.Quiz) error
}

func QuizRepositoryImpl(db *gorm.DB) QuizRepository {
	return &quizRepository{db: db}
}

func (q *quizRepository) Create(quiz *models.Quiz) error {
	return q.db.Create(&quiz).Error
}

func (q *quizRepository) GetQuizByID(QuizID uint) (models.Quiz, error) {
	var dbQuiz models.Quiz
	err := q.db.Where("id = ?", QuizID).First(&dbQuiz).Error
	if err != nil {
		return dbQuiz, err
	}
	return dbQuiz, nil
}

func (q *quizRepository) Delete(quiz *models.Quiz) error {
	return q.db.Delete(&quiz).Error
}

func (q *quizRepository) GetAllQuizzes() ([]models.Quiz, error) {
	var dbQuizzes []models.Quiz
	err := q.db.Find(&dbQuizzes).Error
	if err != nil {
		return dbQuizzes, err
	}
	return dbQuizzes, nil
}

func (q *quizRepository) Update(quiz *models.Quiz) error {
	return q.db.Save(&quiz).Error
}
