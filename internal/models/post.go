package models

import "time"

const postRecordType = "post"

type PostRecord struct {
	Record
	Post
}

type Post struct {
    ID        string    `dynamodbav:"id"        json:"id"`
    Filename  string    `dynamodbav:"filename"  json:"filename"`
    BodyURI   string    `dynamodbav:"bodyURI"   json:"bodyURI"`
    UserID    string    `dynamodbav:"userID"    json:"userID"` 
    CreatedAt time.Time `dynamodbav:"CreatedAt" json:"createdAt"`
}

type PostDto struct {
    Filename string
    BodyURI  string
}

func NewPostRecordPK(userID string) string {
	return NewUserRecordKey(userID)
}

func NewPostRecordSK(postID string) string {
	return postRecordType + "/" + postID
}

