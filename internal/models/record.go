package models

import "time"

type Record struct {
	ID         string `json:"id"`
	Range      string `json:"rng"`
	RecordType string `json:"typ"`
	Version    int    `json:"v"`
}

type UserRecord struct {
	Record
	UserRecordFields
}

type UserRecordFields struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

const userRecordName = "user"

func NewUserRecord(user User) UserRecord {
	var ur UserRecord
	ur.ID = NewUserRecordHashKey(user.ID)
	ur.Range = NewUserRecordRangeKey()
	ur.Email = user.ID
	ur.Username = user.Username
	ur.Password = user.Password
	ur.CreatedAt = user.CreatedAt

	return ur
}

func NewUserRecordHashKey(email string) string {
	return userRecordName + "/" + email
}

func NewUserRecordRangeKey() string {
	return userRecordName
}

