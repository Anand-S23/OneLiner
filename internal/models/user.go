package models

import (
	"time"
)

const userRecordName = "user"

type UserRecord struct {
	Record
    User
}

type User struct {
    ID        string    `dynamodbav:"id"        json:"id"`
    Email     string    `dynamodbav:"email"     json:"email"`
    Password  string    `dynamodbav:"password"  json:"-"`
    CreatedAt time.Time `dynamodbav:"CreatedAt" json:"createdAt"`
}

