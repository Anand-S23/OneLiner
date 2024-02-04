package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

const COOKIE_NAME = "jwt"

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
    token, err := GenerateToken(c.JwtSecretKey, user.ID, expDuration)
    if err != nil {
        log.Println("Error generating token")
        return InternalServerError(w)
    }

    cookie := GenerateCookie(c.CookieSecret, COOKIE_NAME, token, expDuration)
    log.Println(cookie)
    if cookie == nil {
        log.Println("Error generating cookie")
        return InternalServerError(w)
    }
    http.SetCookie(w, cookie)

    successMsg := map[string]string {
        "message": "User logged in successfully",
    }
    log.Println("User successfully logged in")
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) error {
    cookie := GenerateExpiredCookie(COOKIE_NAME)
    http.SetCookie(w, cookie)
    log.Println("User successfully logged out")
    return WriteJSON(w, http.StatusOK, "")
}

