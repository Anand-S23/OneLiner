package models

import "time"

const postRecordName = "post"

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

