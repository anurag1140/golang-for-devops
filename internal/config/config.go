package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	S3Bucket  string
	AWSRegion string

	DatabaseURL string

	// DBHost     string
	// DBPort     string
	// DBUser     string
	// DBPassword string
	// DBName     string
}

func Load() Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	return Config{
		Port: getEnv("PORT", "8080"),

		AWSRegion: getEnv("AWS_REGION", "ap-south-1"),

		S3Bucket: getEnv("S3_BUCKET", ""),

		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
