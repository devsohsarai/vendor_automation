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

func Secret() string {
	return os.Getenv("SECRET")
}

func DbMongo() string {
	return os.Getenv("DB_MONGO")
}
