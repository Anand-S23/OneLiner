package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

func getUserFromResquest(r *http.Request, jwtSecretKey []byte, cookieSecret *securecookie.SecureCookie) (*models.Claims, error) {
    tokenString, err := models.ParseCookie(r, cookieSecret, models.COOKIE_NAME)
	if err != nil {
        errMsg := fmt.Sprintf("Invalid request, could not parse cookie: %s", err)
        return nil, errors.New(errMsg)
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
        errMsg := fmt.Sprintf("Invalid cookie, could not parse token: %s", err)
        return nil, errors.New(errMsg)
	}
    
	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
        return nil, errors.New("Invalid token, not able to parse claims")
	}

    return claims, nil
}

func Authentication(next http.Handler, jwtSecretKey []byte, cookieSecret *securecookie.SecureCookie) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims, err := getUserFromResquest(r, jwtSecretKey, cookieSecret)
        if err != nil {
            log.Println(err.Error())
            controller.UnauthorizedError(w)
            return
        }
        
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
    })
}

