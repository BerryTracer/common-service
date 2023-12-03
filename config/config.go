package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(envName string) (string, error) {
	env := os.Getenv(envName)
	if env == "" {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			return "", err
		}

		env = os.Getenv(envName)
		if env == "" {
			return "", errors.New(envName + " environment variable not set")
		}
	}

	return env, nil
}
