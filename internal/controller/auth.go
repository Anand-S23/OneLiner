package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/validators"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func createToken(secretKey string, userID string, expDuration time.Duration) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(expDuration).Unix()

    return token.SignedString([]byte(secretKey))
}

func createJWTCookie(token string, expDuration time.Duration, secure bool) http.Cookie {
    return http.Cookie {
        Name: "jwt",
        Value: token,
        Expires: time.Now().Add(expDuration),
        HttpOnly: true,
        Secure: secure,
        Path: "/",
        SameSite: http.SameSiteStrictMode,
    }
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) error {
    var userData models.RegisterDto
    err := json.NewDecoder(r.Body).Decode(&userData)
    if err != nil {
        errMsg := map[string]string {
            "error": "Could not parse sign up data",
        }
        return WriteJSON(w, http.StatusBadRequest, errMsg)
    }

    authErrs := validators.AuthValidator(userData, c.store)
    if len(authErrs) != 0 {
        return WriteJSON(w, http.StatusBadRequest, authErrs)
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
        log.Println("Error hashing the password")
        return InternalServerError(w)
	}
    userData.Password = string(hashedPassword)

    user := models.NewUser(userData)
    err = c.store.PutUser(user)
    if err != nil {
        log.Println("Error storing the password in the database")
        return InternalServerError(w)
    }

    expDuration := time.Hour * 24
    token, err := createToken(c.JwtSecretKey, user.ID, expDuration)
    if err != nil {
        log.Println("Error creating token")
        return InternalServerError(w)
    }

    cookie := createJWTCookie(token, expDuration, c.production)
    http.SetCookie(w, &cookie)

    successMsg := map[string]string {
        "message": "User created successfully",
        "userID": user.ID,
    }

    return WriteJSON(w, http.StatusOK, successMsg)
}

