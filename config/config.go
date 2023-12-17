package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type EnvLoader interface {
	Getenv(key string) string
	Load() error
}

type RealEnvLoader struct{}

func NewRealEnvLoader() RealEnvLoader {
	return RealEnvLoader{}
}

func (RealEnvLoader) Getenv(key string) string {
	return os.Getenv(key)
}

func (RealEnvLoader) Load() error {
	return godotenv.Load()
}

func LoadEnv(envLoader EnvLoader, envName string) (string, error) {
	env := envLoader.Getenv(envName)
	if env == "" {
		// Load .env file
		err := envLoader.Load()
		if err != nil {
			return "", err
		}

		env = envLoader.Getenv(envName)
		if env == "" {
			return "", errors.New(envName + " environment variable not set")
		}
	}

	return env, nil
}

func LoadEnvWithDefault(envLoader EnvLoader, envName string, defaultValue string) (string, error) {
	env := envLoader.Getenv(envName)
	if env == "" {
		// Load .env file
		err := envLoader.Load()
		if err != nil {
			// If .env loading fails, use the default value
			if defaultValue != "" {
				return defaultValue, nil
			}
			return "", err
		}

		env = envLoader.Getenv(envName)
		if env == "" {
			if defaultValue != "" {
				return defaultValue, nil
			}
			return "", errors.New(envName + " environment variable not set")
		}
	}

	return env, nil
}
