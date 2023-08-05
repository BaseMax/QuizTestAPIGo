package routers

import (
	"net/http"

	"github.com/BaseMax/QuizTestAPIGo/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RedirectToDocs(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}

func SetupRoutes(db *gorm.DB, gin *gin.Engine) {
	swaggerRoute := gin.Group("")
	docs.SwaggerInfo.BasePath = "/api/"
	swaggerRoute.GET("/", RedirectToDocs)
	swaggerRoute.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
