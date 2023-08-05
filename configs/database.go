package configs

import (
	"log"
	"os"

	"github.com/BaseMax/QuizTestAPIGo/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	log.Println("Mysql is Connected Successfuly!")
	DB.AutoMigrate(&models.User{}, &models.Question{}, &models.Quiz{}, &models.QuizResult{}, &models.QuizResultDetail{})
	log.Println("Tables migrations was successfully")
	return DB, nil
}
