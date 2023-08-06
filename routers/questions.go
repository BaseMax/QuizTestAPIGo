package routers

import (
	"github.com/BaseMax/QuizTestAPIGo/controllers"
	"github.com/BaseMax/QuizTestAPIGo/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuestionCRUDRouter(db *gorm.DB, r *gin.RouterGroup) {
	questionRepo := repository.QuestionRepositoryImpl(db)
	questionController := &controllers.QuestionController{Service: questionRepo}
	r.GET("/:id/questions", questionController.GetAllQuestionsHandler)
	r.POST("/:id/questions", questionController.CreateQuestionHandler)
	r.PUT("/:id/questions/:question_id", questionController.UpdateQuestionHandler)	
	r.DELETE("/:id/questions/:question_id", questionController.DeleteQuestionHandler)
}
