package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using default environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}

	return &Config{
		Port: port,
	}
}
