// main.go
package main

import (
	"github.com/joho/godotenv"
	"log"
	"yt-cli/app/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
	cmd.Execute()
}
