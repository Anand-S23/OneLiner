package models

import (
	"time"
)

type User struct {
    ID        string // ID is an email address
    Password  string
    CreatedAt time.Time
}

type UserDetail struct {
    User
    Posts []Post
}

func NewUser(userData RegisterDto) User {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
    return User {
        ID: userData.Email,
        Password: userData.Password,
        CreatedAt: now,
    }
}

func NewUserFromRecord(ur UserRecord) *User {
    return &User {
        ID: ur.Email,
        Password: ur.Password,
        CreatedAt: ur.CreatedAt,
    }
}

