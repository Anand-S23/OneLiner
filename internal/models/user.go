package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

const userRecordType = "user"

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

func NewUserUUID(email string) uuid.UUID {
	hasher := md5.New()
	hasher.Write([]byte(email))
	hash := hex.EncodeToString(hasher.Sum(nil))

	uuidFromHash := uuid.NewSHA1(uuid.Nil, []byte(hash))
	return uuidFromHash
}

func NewUser(userData RegisterDto) User {
    id := NewUserUUID(userData.Email).String()
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return User {
        ID: id,
        Email: userData.Email,
        Password: userData.Password,
        CreatedAt: now,
    }
}

func NewUserRecord(user User) UserRecord {
    var ur UserRecord
    ur.PK = NewUserRecordPK(user.ID)
    ur.SK = NewUserRecordSK(user.Email)
    ur.Type = userRecordType
    ur.ID = user.ID
    ur.Email = user.Email
    ur.Password = user.Password
    ur.CreatedAt = user.CreatedAt

    return ur
}

func NewUserRecordPK(id string) string {
	return userRecordType + "/" + id
}

func NewUserRecordSK(email string) string {
	return userRecordType + "/" + email
}

