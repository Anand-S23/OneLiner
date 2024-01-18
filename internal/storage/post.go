package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
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
