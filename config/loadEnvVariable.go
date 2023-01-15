package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariable() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
