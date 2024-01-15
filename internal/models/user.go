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
    id := NewUserUUID(email).String()
    pk := NewUserRecordKey(id)
    sk := NewUserRecordKey(email)
    return pk, sk
}

func NewUserRecordKey(id string) string {
	return userRecordType + "/" + id
}

