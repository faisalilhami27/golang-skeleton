package util

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)

func LoadEnv() error {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	return env
}

func ConvertToIndonesiaTimezone() time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	result := time.Now().In(location)
	return result
}
