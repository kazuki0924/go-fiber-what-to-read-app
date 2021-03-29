package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "prod" {
		err := godotenv.Load(".env.prod")
		if err != nil {
			fmt.Println("Error loading .env.prod file")
			panic(err)
		}
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
}
