package config

import "os"

var AppConfig Environment

type Environment struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBDatabase string
	Port       string
}

func InitApp() {
	AppConfig.DBDatabase = os.Getenv("DB_DATABASE")
	if AppConfig.DBDatabase == "" {
		panic("DB_DATABASE IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBHost = os.Getenv("DB_HOST")
	if AppConfig.DBHost == "" {
		panic("DB_HOST IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBPort = os.Getenv("DB_PORT")
	if AppConfig.DBPort == "" {
		panic("DB_PORT IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBUsername = os.Getenv("DB_USERNAME")
	if AppConfig.DBUsername == "" {
		panic("DB_USERNAME IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.DBPassword = os.Getenv("DB_PASSWORD")
	if AppConfig.DBPassword == "" {
		panic("DB_PASSWORD IS EMPTY, PLEASE CONFIGURE FIRST")
	}

	AppConfig.Port = os.Getenv("PORT")
	if AppConfig.Port == "" {
		panic("PORT IS EMPTY, PLEASE CONFIGURE FIRST")
	}
}
