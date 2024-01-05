package models

import "time"

type Post struct {
    ID        string
    Name      string
    BodyURI   string
    UserID    string
    CreatedAt time.Time
}

