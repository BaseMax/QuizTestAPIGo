package controllers

import (
	"net/http"
	"strconv"

	"github.com/BaseMax/QuizTestAPIGo/entity"
	"github.com/BaseMax/QuizTestAPIGo/models"
	"github.com/BaseMax/QuizTestAPIGo/repository"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	Service repository.QuestionRepository
}

func (q *QuestionController) CreateQuestionHandler(ctx *gin.Context) {
	var data entity.QuestionRequest
	quizID, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if _, err := q.Service.GetQuizByID(uint(quizID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz with this id not found"})
		return
	}

	newQuestion := models.Question{Text: data.Text, QuizID: uint(quizID)}
	if err := q.Service.Create(&newQuestion); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newQuestion)
}

func (q *QuestionController) DeleteQuestionHandler(ctx *gin.Context) {
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	questionID, _ := strconv.Atoi(ctx.Param("question_id"))

	if _, err := q.Service.GetQuizByID(uint(quizID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz with this id not found"})
		return
	}

	question, err := q.Service.GetQuestionByID(uint(questionID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "question with this id not found"})
		return
	}

	if err := q.Service.Delete(&question); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Delete question record falied"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (q *QuestionController) GetAllQuestionsHandler(ctx *gin.Context) {
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	questions, err := q.Service.GetAllQuestions(uint(quizID))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "fetch questions falied"})
		return
	}

	ctx.JSON(http.StatusOK, questions)
}

func (q *QuestionController) UpdateQuestionHandler(ctx *gin.Context) {
	var data entity.QuestionRequest
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	questionID, _ := strconv.Atoi(ctx.Param("question_id"))

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	if _, err := q.Service.GetQuizByID(uint(quizID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz with this id not found"})
		return
	}

	question, err := q.Service.GetQuestionByID(uint(questionID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "question with this id not found"})
		return
	}

	question.Text = data.Text
	if err := q.Service.Update(&question); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "update question record falied"})
		return
	}

	ctx.JSON(http.StatusOK, question)
}
