package main

import (
	"fmt"
	"os"

	infrastructure "github.com/kazuki0924/go-what-to-read-app/infrastructure/env"
)

func main() {
	infrastructure.LoadEnv()

	port := os.Getenv("HTTP_PORT")

	fmt.Println(port)
}
