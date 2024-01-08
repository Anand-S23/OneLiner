package validators

import (
	"errors"
	"net/mail"
	"regexp"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/storage"
)


func AuthValidator(userData models.RegisterDto, store *storage.DynamoStore) map[string]string {
    errs := make(map[string]string, 3)

    err := validateUsername(userData.Username, store)
    if err != nil {
        errs["username"] = err.Error()
    }

    err = validateEmail(userData.Email, store)
    if err != nil {
        errs["email"] = err.Error()
    }

    err = validatePassword(userData.Password, userData.Confirm)
    if err != nil {
        errs["password"] = err.Error()
    }

    return errs
}

func validateUsername(username string, store *storage.DynamoStore) error {
    if len(username) < 5 || len(username) > 25 {
        return errors.New("Username must be between 5-15 characters long")
    }

    if !regexp.MustCompile("^[a-zA-Z0-9_\\-.]+$").MatchString(username) {
        return errors.New("Username can only contain letters, numbers, hyphens(-), underscores(_) and periods(.)")
    }

    // TODO: Unique Username

    return nil
}

func validateEmail(email string, store *storage.DynamoStore) error {
    _, err := mail.ParseAddress(email)
    if err != nil {
        return errors.New("Email entered is not valid")
    }

    if len(email) > 64 {
        return errors.New("Email must be less than 64 characters")
    }

    // TODO: Unique Email

    return nil
}

func validatePassword(password string, confirm string) error {
    if len(password) < 8 || len(password) > 30 {
        return errors.New("Password must be between 8 and 30 characters long")
    }

    if password != confirm {
        return errors.New("Passwords must match")
    }

    return nil
}

