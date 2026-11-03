package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}

	return val, nil
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, fmt.Errorf("error getting env MONGO_URI: %w", err)
	}

	mongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, fmt.Errorf("error getting env MONGO_DB_NAME: %w", err)
	}

	serverPort, err := extractEnv("SERVER_PORT")
	if err != nil {
		return Config{}, fmt.Errorf("error getting env SERVER_PORT: %w", err)
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: serverPort,
	}, nil
}
