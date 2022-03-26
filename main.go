package main

import (
	"log"
	"triadmoko-be-golang/config"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.Database()
}
