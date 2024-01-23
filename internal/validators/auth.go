package validators

import (
	"errors"
	"net/mail"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/Anand-S23/Snippet/internal/storage"
)


func AuthValidator(userData models.RegisterDto, store *storage.SnippetStore) map[string]string {
    errs := make(map[string]string, 2)

    err := validateEmail(userData.Email, store)
    if err != nil {
        errs["email"] = err.Error()
    }

    err = validatePassword(userData.Password, userData.Confirm)
    if err != nil {
        errs["password"] = err.Error()
    }

    return errs
}

func validateEmail(email string, store *storage.SnippetStore) error {
    _, err := mail.ParseAddress(email)
    if err != nil {
        return errors.New("Email entered is not valid")
    }

    if len(email) > 64 {
        return errors.New("Email must be less than 64 characters")
    }

    user := store.GetUser(models.GetKeysFromEmail(email))
    if user.ID != "" {
        return errors.New("User already exsits with that email")
    }

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

