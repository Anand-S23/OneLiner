package storage

import (
	"context"
	"io"
	"log"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (store *SnippetStore) UploadFileToS3(file io.Reader, key string) error {
    _, err := store.s3.Client.PutObject(context.TODO(), &s3.PutObjectInput {
		Bucket: store.s3.BucketName,
		Key:    aws.String(key),
		Body:   file,
	})

    return err
}

func (store *SnippetStore) PutPost(post models.Post) error {
    pr := models.NewPostRecord(post)
    item, err := attributevalue.MarshalMap(pr)
    if err != nil {
        log.Println("Could not marshal item")
        return err
    }

    putItemInput := &dynamodb.PutItemInput {
        TableName: store.tableName,
        Item: item,
    }

    _, err = store.db.PutItem(context.TODO(), putItemInput)
    if err != nil {
        log.Println("Could not put item in db")
        return err
    }

    return nil
}

