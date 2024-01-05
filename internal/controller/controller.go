package controller

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/storage"
)

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

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

