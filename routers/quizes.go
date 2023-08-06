package routers

import (
	"github.com/BaseMax/QuizTestAPIGo/controllers"
	"github.com/BaseMax/QuizTestAPIGo/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuizCRUDRouter(db *gorm.DB, r *gin.RouterGroup) {
	quizRepo := repository.QuizRepositoryImpl(db)
	quizController := &controllers.QuizController{Service: quizRepo}
	r.GET("/", quizController.GetAllQuizzesHandler)
	r.GET("/:id", quizController.GetQuizByIDHandler)
	r.POST("/", quizController.CreateQuizHandler)
	r.PUT("/:id", quizController.UpdateQuizHandler)
	r.DELETE("/:id", quizController.DeleteQuizHandler)
}
