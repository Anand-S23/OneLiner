package models

import (
	"time"
)

type Record struct {
	ID         string `dynamodbav:"id" json:"id"`
	Range      string `dynamodbav:"rng" json:"rng"`
	RecordType string `dynamodbav:"typ" json:"typ"`
	Version    int    `dynamodbav:"v" json:"v"`
}

type UserRecord struct {
	Record
	UserRecordFields
}

type UserRecordFields struct {
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

const userRecordName = "user"

func NewUserRecord(user User) UserRecord {
	var ur UserRecord
	ur.ID = NewUserRecordHashKey(user.ID)
	ur.Range = NewUserRecordRangeKey()
	ur.Email = user.ID
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

type PostRecord struct {
	Record
	PostRecordFields
}

type PostRecordFields struct {
    Name      string    `json:"name"`
    BodyURI   string    `json:"bodyURI"`
    UserID    string    `json:"userID"`
    CreatedAt time.Time `json:"createdAt"`
}

const postRecordName = "post"

func NewPostRecord(post Post) PostRecord {
	var pr PostRecord
	pr.ID = NewPostRecordHashKey(post.ID)
	pr.Range = NewPostRecordRangeKey()
	pr.Name = post.Name
	pr.BodyURI = post.BodyURI
	pr.UserID = post.UserID
	pr.CreatedAt = post.CreatedAt

	return pr
}

func NewPostRecordHashKey(id string) string {
	return postRecordName + "/" + id
}

func NewPostRecordRangeKey() string {
	return postRecordName
}
