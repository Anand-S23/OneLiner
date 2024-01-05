package models

import "time"

type User struct {
    ID        string // ID is an email address
    Username  string
    Password  string
    CreatedAt time.Time
}

type UserDetail struct {
    User
    Posts []Post
}

func NewUserFromRecord(ur UserRecord) *User {
    return &User {
        ID: ur.Email,
        Username: ur.Username,
        Password: ur.Password,
        CreatedAt: ur.CreatedAt,
    }
}

