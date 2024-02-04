package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

func getUserFromResquest(r *http.Request, jwtSecretKey string, cookieSecret *securecookie.SecureCookie) (*controller.Claims, error) {
    tokenString, err := controller.ParseCookie(r, cookieSecret, controller.COOKIE_NAME)
	if err != nil {
        log.Println(err)
        return nil, errors.New("Invalid request, could not parse cookie")
	}

	token, err := jwt.ParseWithClaims(tokenString, &controller.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
        log.Println(err)
        return nil, errors.New("Invalid cookie, not able to parse token")
	}
    
	claims, ok := token.Claims.(*controller.Claims)
	if !ok || !token.Valid {
        return nil, errors.New("Invalid token, not able to parse claims")
	}

    return claims, nil
}

func Authentication(next http.Handler, jwtSecretKey string, cookieSecret *securecookie.SecureCookie) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims, err := getUserFromResquest(r, jwtSecretKey, cookieSecret)
        if err != nil {
            // TODO: Need to figure out if UnauthorizedServerError can be used here
            errMsg := controller.ErrorMessage("Unauthorized")
            controller.WriteJSON(w, http.StatusUnauthorized, errMsg)
            log.Println(err.Error())

            // TODO: redirect
            return
        }
        
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
    })
}

