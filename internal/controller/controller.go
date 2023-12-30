package controller

import "github.com/Anand-S23/OneLiner/internal/storage"

type Controller struct {
    store        *storage.DynamoStore
    production   bool
    JwtSecretKey string
}

func NewController(store *storage.DynamoStore, secretKey string, production bool) *Controller {
    return &Controller {
        store: store,
        production: production,
        JwtSecretKey: secretKey,
    }
}
