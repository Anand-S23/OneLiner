package config

import (
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION bool
    PORT       string
}

func LoadEnv() (*EnvVars, error) {
    envMode := GetEnv("MODE", "development")
    port := GetEnv("PORT", "8080")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        PORT: port,
    }, nil
}

func GetEnv(env, defaultValue string) string {
	variable := os.Getenv(env)
	if variable == "" {
		return defaultValue
	}

	return variable
}

func GetEnvOrPanic(env string, message string) string {
	variable := os.Getenv(env)
	if variable == "" {
        log.Fatal(message)
	}

	return variable
} 

