package controller

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/storage"
	"github.com/gorilla/securecookie"
)

type Controller struct {
    store        *storage.SnippetStore
    production   bool
    JwtSecretKey []byte
    CookieSecret *securecookie.SecureCookie
}

func NewController(store *storage.SnippetStore, secretKey []byte, cookieHashKey []byte, cookieBlockKey []byte, production bool) *Controller {
    return &Controller {
        store: store,
        production: production,
        JwtSecretKey: secretKey,
        CookieSecret: securecookie.New(cookieHashKey, cookieBlockKey),
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

