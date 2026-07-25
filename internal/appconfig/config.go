package config

import "os"

type Config struct {
	Port      string
	S3Bucket  string
	AWSRegion string
}

// cfg:=config.Load()
func Load() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		AWSRegion: getEnv("AWS_REGION", "ap-south-1"),
		S3Bucket:  getEnv("S3_BUCKET", ""),
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
