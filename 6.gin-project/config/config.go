package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found — using system environment variables")
	}

	JwtSecret = os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		log.Fatalf("JWT_SECRET is not set in the environment variables")
	}
}
