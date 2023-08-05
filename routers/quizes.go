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
	r.POST("/", quizController.CreateQuizHandler)
	r.GET("/:id", quizController.GetQuizByIDHandler)
	r.GET("/", quizController.GetAllQuizzesHandler)
	r.DELETE("/:id", quizController.DeleteQuizHandler)
	r.PUT("/:id", quizController.UpdateQuizHandler)
}
