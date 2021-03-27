package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "prod" {
		err := godotenv.Load(".env.prod")
		if err != nil {
			log.Fatal("Error loading .env.prod file")
		}
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	LoadEnv()
}
