package config

import (
	"log"
	"os"
)

type EnvVars struct {
    PRODUCTION       bool
    JWT_SECRET       string
    COOKIE_HASH_KEY  string
    COOKIE_BLOCK_KEY string
    PORT             string
    S3_BUCKET        string
}

func LoadEnv() (*EnvVars, error) {
    envMode := GetEnv("MODE", "development")
    secret := GetEnvOrPanic("JWT_SECRET", "Must provide JWT_SECRET variable in .env file")
    hashKey := GetEnvOrPanic("COOKIE_HASH_KEY", "Must provide COOKIE_HASH_KEY variable in .env file")
    blockKey := GetEnvOrPanic("COOKIE_BLOCK_KEY", "Must provide COOKIE_BLOCK_KEY variable in .env file")
    port := GetEnv("PORT", "8080")
    s3Bucket := GetEnv("S3_BUCKET", "s3bucket-snippet")

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        JWT_SECRET: secret,
        COOKIE_HASH_KEY: hashKey,
        COOKIE_BLOCK_KEY: blockKey,
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

