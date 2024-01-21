package models

import (
	"time"
)

const userRecordType = "user"

type UserRecord struct {
	Record
    User
}

type User struct {
    ID        string    `dynamodbav:"ID"        json:"id"`
    Email     string    `dynamodbav:"Email"     json:"email"`
    Password  string    `dynamodbav:"Password"  json:"-"`
    CreatedAt time.Time `dynamodbav:"CreatedAt" json:"createdAt"`
}

func NewUser(userData RegisterDto) User {
    id := NewHashedUUID(userData.Email)
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return User {
        ID: id,
        Email: userData.Email,
        Password: userData.Password,
        CreatedAt: now,
    }
}

func NewUserFromRecord(ur UserRecord) User {
    return User {
        ID: ur.ID,
        Email: ur.Email,
        Password: ur.Password,
        CreatedAt: ur.CreatedAt,
    }
}

func NewUserRecord(user User) UserRecord {
    var ur UserRecord
    ur.PK = NewUserRecordKey(user.ID)
    ur.SK = NewUserRecordKey(user.Email)
    ur.Type = userRecordType
    ur.ID = user.ID
    ur.Email = user.Email
    ur.Password = user.Password
    ur.CreatedAt = user.CreatedAt

    return ur
}

func GetKeysFromEmail(email string) (string, string) {
    id := NewHashedUUID(email)
    pk := NewUserRecordKey(id)
    sk := NewUserRecordKey(email)
    return pk, sk
}

func NewUserRecordKey(id string) string {
	return userRecordType + "/" + id
}

