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

func createExpiredJWTCookie(secure bool) http.Cookie {
    return http.Cookie {
        Name: "jwt",
        Value: "",
        Expires: time.Unix(0, 0),
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
        return BadRequestError(w, "Could not parse sign up data")
    }

    authErrs := validators.AuthValidator(userData, c.store)
    if len(authErrs) != 0 {
        log.Println("Failed to create new user, invalid data")
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
        log.Printf("Error storing the password in the database, %s", err)
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

    log.Println("User created successfully")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) error {
    var loginData models.LoginDto
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        return BadRequestError(w, "Could not parse login data")
    }

    user := c.store.GetUser(models.GetKeysFromEmail(loginData.Email))
    if user.ID == "" {
        return BadRequestError(w, "Incorrect email or password, please try again")
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
        return BadRequestError(w, "Incorrect email or password, please try again")
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
        "message": "User logged in successfully",
    }
    log.Println("User successfully logged in")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) error {
    cookie := createExpiredJWTCookie(c.production)
    http.SetCookie(w, &cookie)
    log.Println("User successfully logged out")
    return WriteJSON(w, http.StatusOK, "")
}

