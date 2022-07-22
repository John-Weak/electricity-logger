package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	print("Loaded .env file \n")
}

func EnvMongoURI() string {
	return os.Getenv("MONGODB_URI")
}
