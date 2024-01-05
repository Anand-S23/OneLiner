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

