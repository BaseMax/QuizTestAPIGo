package main

import (
	"github.com/BaseMax/QuizTestAPIGo/configs"
	"github.com/BaseMax/QuizTestAPIGo/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
	configs.GetDB()
}

func main() {
	srv := gin.Default()
	srv.Use(gin.Logger())
	srv.Use(gin.Recovery())
	routers.SetupRoutes(configs.DB, srv)
	srv.Run(":8000")
}
