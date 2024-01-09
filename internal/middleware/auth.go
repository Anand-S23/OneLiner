package middleware

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/dgrijalva/jwt-go"
)

func Authentication(next http.Handler, jwtSecretKey string) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        errMsg := map[string]string {"error": "Unauthorized",}

        cookie, err := r.Cookie("jwt")
        if err != nil || cookie.Value == "" {
            controller.WriteJSON(w, http.StatusUnauthorized, errMsg)
            return
        }

        token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecretKey), nil
        })

        if err != nil || !token.Valid {
            controller.WriteJSON(w, http.StatusUnauthorized, errMsg)
            return
        }

        next.ServeHTTP(w, r)
    })
}

