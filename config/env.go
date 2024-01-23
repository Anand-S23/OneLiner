package config

import (
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION bool
    JWT_SECRET string
    PORT       string
    S3_BUCKET  string
}

func LoadEnv() (*EnvVars, error) {
    envMode := GetEnv("MODE", "development")
    secret := GetEnvOrPanic("JWT_SECRET", "Must provide JWT_SECRET variable in .env file")
    port := GetEnv("PORT", "8080")
    s3Bucket := GetEnv("S3_BUCKET", "s3bucket-snippet")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        JWT_SECRET: secret,
        PORT: port,
        S3_BUCKET: s3Bucket,
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

