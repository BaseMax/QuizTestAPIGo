package controllers

import (
	"net/http"
	"strconv"

	"github.com/BaseMax/QuizTestAPIGo/entity"
	"github.com/BaseMax/QuizTestAPIGo/models"
	"github.com/BaseMax/QuizTestAPIGo/repository"
	"github.com/gin-gonic/gin"
)

type QuizController struct {
	Service repository.QuizRepository
}

func (q *QuizController) CreateQuizHandler(ctx *gin.Context) {
	var data entity.QuizRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	newQuiz := models.Quiz{Title: data.Title, Description: data.Description}
	if err := q.Service.Create(&newQuiz); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newQuiz)
}

func (q *QuizController) GetQuizByIDHandler(ctx *gin.Context) {
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	quiz, err := q.Service.GetQuizByID(uint(quizID))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz not found"})
		return
	}

	ctx.JSON(http.StatusOK, quiz)
}

func (q *QuizController) DeleteQuizHandler(ctx *gin.Context) {
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	quiz, err := q.Service.GetQuizByID(uint(quizID))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz not found"})
		return
	}

	if err := q.Service.Delete(&quiz); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "Delete quiz record falied"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (q *QuizController) GetAllQuizzesHandler(ctx *gin.Context) {
	quizzes, err := q.Service.GetAllQuizzes()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": "fetch quizzes falied"})
		return
	}

	ctx.JSON(http.StatusOK, quizzes)
}

func (q *QuizController) UpdateQuizHandler(ctx *gin.Context) {
	var data entity.QuizRequest
	quizID, _ := strconv.Atoi(ctx.Param("id"))
	quiz, err := q.Service.GetQuizByID(uint(quizID))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": "quiz not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		return
	}

	quiz.Title = data.Title
	quiz.Description = data.Description
	if err := q.Service.Update(&quiz); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "updated quiz failed"})
		return
	}
	ctx.JSON(http.StatusOK, quiz)
}
