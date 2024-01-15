package storage

import (
	"github.com/Anand-S23/Snippet/internal/blob"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoStore struct {
	db        *dynamodb.Client
    tableName *string
    s3        *blob.S3Bucket
}

func NewDynamoStore(db *dynamodb.Client, tableName string, s3 *blob.S3Bucket) *DynamoStore {
    return &DynamoStore{
        db: db,
        tableName: aws.String(tableName),
        s3: s3,
    }
}

