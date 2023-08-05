package main

import (
	"github.com/BaseMax/QuizTestAPIGo/configs"
	"github.com/gin-gonic/gin"
)


func init(){
	configs.LoadEnv()
	configs.GetDB()
}

func main() {
	srv := gin.Default()
	srv.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"response": "Hi"}) })
	srv.Run(":8000")
}
