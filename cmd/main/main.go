package main

import (
	"log"

	"github.com/epousa/ginFrameworkPractise/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server.StartServer()
}
