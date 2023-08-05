package configs

import "github.com/joho/godotenv"

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}
