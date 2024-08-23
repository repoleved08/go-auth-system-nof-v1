package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load environment variables: err=%v", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
