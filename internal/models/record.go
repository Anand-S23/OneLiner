package models

import "time"

type record struct {
	ID         string `json:"id"`
	Range      string `json:"rng"`
	RecordType string `json:"typ"`
	Version    int    `json:"v"`
}

type userRecord struct {
	record
	userRecordFields
}

type userRecordFields struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

const userRecordName = "user"

func newUserRecord(user User) userRecord {
	var ur userRecord
	ur.ID = newUserRecordHashKey(user.ID)
	ur.Range = newUserRecordRangeKey()
	ur.Email = user.ID
	ur.Username = user.Username
	ur.Password = user.Password
	ur.CreatedAt = user.CreatedAt

	return ur
}

func newUserRecordHashKey(email string) string {
	return userRecordName + "/" + email
}

func newUserRecordRangeKey() string {
	return userRecordName
}

