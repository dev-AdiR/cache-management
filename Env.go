package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	ClientUrl string
}

func LoadEnv() *Environment {
	_ = godotenv.Load()

	return &Environment{
		ClientUrl: os.Getenv("Client_URL"),
	}
}
