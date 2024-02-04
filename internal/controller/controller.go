package controller

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/storage"
	"github.com/gorilla/securecookie"
)

type Controller struct {
    store        *storage.SnippetStore
    production   bool
    JwtSecretKey string
    CookieSecret *securecookie.SecureCookie
}

func NewController(store *storage.SnippetStore, secretKey string, cookieHashKey string, cookieBlockKey string, production bool) *Controller {
    return &Controller {
        store: store,
        production: production,
        JwtSecretKey: secretKey,
        CookieSecret: securecookie.New([]byte(cookieHashKey), []byte(cookieBlockKey)),
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

