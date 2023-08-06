package repository

import (
	"github.com/BaseMax/QuizTestAPIGo/models"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}
type QuestionRepository interface {
	Create(q *models.Question) error
	GetQuizByID(quizID uint) (models.Quiz, error)
	GetQuestionByID(questionID uint) (models.Question, error)
	GetAllQuestions(quizID uint) ([]models.Question, error)
	Update(quiz *models.Question) error
	Delete(quiz *models.Question) error
}

func QuestionRepositoryImpl(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (q *questionRepository) Create(question *models.Question) error {
	return q.db.Create(&question).Error
}

func (q *questionRepository) GetQuizByID(quizID uint) (models.Quiz, error) {
	var dbQuiz models.Quiz
	err := q.db.Where("id = ?", quizID).First(&dbQuiz).Error
	if err != nil {
		return dbQuiz, err
	}
	return dbQuiz, nil
}

func (q *questionRepository) GetQuestionByID(questionID uint) (models.Question, error) {
	var dbQuestion models.Question
	err := q.db.Where("id = ?", questionID).First(&dbQuestion).Error
	if err != nil {
		return dbQuestion, err
	}
	return dbQuestion, nil
}

func (q *questionRepository) Delete(question *models.Question) error {
	return q.db.Delete(&question).Error
}

func (q *questionRepository) Update(question *models.Question) error {
	return q.db.Save(&question).Error
}

func (q *questionRepository) GetAllQuestions(quizID uint) ([]models.Question, error) {
	var dbQuestions []models.Question
	err := q.db.Where("quiz_id = ?", quizID).Find(&dbQuestions).Error
	if err != nil {
		return dbQuestions, err
	}
	return dbQuestions, nil
}
