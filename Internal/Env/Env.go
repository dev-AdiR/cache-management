package env

import (
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	ClientUrl string
}

func LoadEnv() *environment {
	_ = godotenv.Load()

	return &environment{
		ClientUrl: os.Getenv("Client_URL"),
	}
}
