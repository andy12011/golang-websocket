package config

import (
	"log"

	"github.com/joho/godotenv")


func Init(path string) {
	err := godotenv.Load(path)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}