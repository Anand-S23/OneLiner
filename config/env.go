package config

import (
	"fmt"
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION         bool
    JWT_SECRET         []byte
    COOKIE_HASH_KEY    []byte
    COOKIE_BLOCK_KEY   []byte
    PORT               string
    S3_BUCKET          string
}

func LoadEnv() (*EnvVars, error) {
    envMode  := GetEnv("MODE", "development")
    port     := GetEnv("PORT", "8080")
    s3Bucket := GetEnv("S3_BUCKET", "s3bucket-snippet")

    secret   := GetEnvOrPanic("JWT_SECRET")
    hashKey  := GetEnvOrPanic("COOKIE_HASH_KEY")
    blockKey := GetEnvOrPanic("COOKIE_BLOCK_KEY")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        JWT_SECRET: []byte(secret),
        COOKIE_HASH_KEY: []byte(hashKey),
        COOKIE_BLOCK_KEY: []byte(blockKey),
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

func GetEnvOrPanic(env string) string {
	variable := os.Getenv(env)
	if variable == "" {
        message := fmt.Sprintf("Must provide %s variable in .env file", env)
        log.Fatal(message)
	}

	return variable
} 

