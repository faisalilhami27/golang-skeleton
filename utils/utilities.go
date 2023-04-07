package util

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() error {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	return env
}
