package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return os.Getenv("MONGOURI")
}
