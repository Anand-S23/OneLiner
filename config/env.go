package config

import (
    "os"
)

type EnvVars struct {
    PRODUCTION bool
    PORT       string
}

func LoadEnv() (*EnvVars, error) {
    envMode := os.Getenv("MODE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        PORT: port,
    }, nil
}
