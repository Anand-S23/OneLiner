package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/dgrijalva/jwt-go"
)

func getUserFromResquest(r *http.Request, jwtSecretKey string) (string, error) {
    cookie, err := r.Cookie("jwt")
    if err != nil || cookie.Value == "" {
        return "", errors.New("Invalid request, could not parse jwt cookie")
    }

    token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtSecretKey), nil
    })
    if err != nil || !token.Valid {
        return "", errors.New("Invalid cookie, not able to parse token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("Invalid token, not able to parse claims")
    }

    userID := claims["user_id"].(string)
    return userID, nil
}

func Authentication(next http.Handler, jwtSecretKey string) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID, err := getUserFromResquest(r, jwtSecretKey)
        if err != nil {
            errMsg := map[string]string {"error": "Unauthorized"}
            controller.WriteJSON(w, http.StatusUnauthorized, errMsg)
            log.Println(err.Error())
            // TODO: redirect
            return
        }
        
        ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
    })
}

