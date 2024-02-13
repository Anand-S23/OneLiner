package models

import (
	"time"
)

const PostRecordType = "post"

type PostRecord struct {
	Record
	Post
}

type Post struct {
    ID          string            `dynamodbav:"ID"          json:"id"`
    Name        string            `dynamodbav:"Name"        json:"name"`
    Description string            `dynamodbav:"Description" json:"description"` 
    Files       map[string]string `dynamodbav:"Files"       json:"files"`
    UserID      string            `dynamodbav:"UserID"      json:"userID"` 
    CreatedAt   time.Time         `dynamodbav:"CreatedAt"   json:"createdAt"`
}

type PostDto struct {
    Name        string
    Description string
    Files       map[string]string
}

type DeleteFilesDto struct {
    UserID string
    Files  map[string]string
}

func NewPost(postData PostDto, userID string) Post {
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return Post {
        ID: NewUUID(),
        Name: postData.Name,
        Description: postData.Description,
        Files: postData.Files,
        UserID: userID,
        CreatedAt: now,
    }
}

func NewPostFromRecord(pr PostRecord) Post {
    return Post {
        ID: pr.ID,
        Name: pr.Name,
        Description: pr.Description,
        Files: pr.Files,
        UserID: pr.UserID,
        CreatedAt: pr.CreatedAt,
    }
}

func NewPostRecord(post Post) PostRecord {
    var pr PostRecord
    pr.PK = NewPostRecordPK(post.UserID)
    pr.SK = NewPostRecordSK(post.ID)
    pr.Type = PostRecordType
    pr.ID = post.ID
    pr.Name = post.Name
    pr.Description = post.Description
    pr.Files = post.Files
    pr.UserID = post.UserID
    pr.CreatedAt = post.CreatedAt

    return pr
}

func NewPostRecordPK(userID string) string {
	return NewUserRecordKey(userID)
}

func NewPostRecordSK(postID string) string {
	return PostRecordType + "/" + postID
}

