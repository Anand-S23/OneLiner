package storage

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoStore struct {
	db        *dynamodb.Client
    tableName *string
    now       func() time.Time
}

func NewDynamoStore(db *dynamodb.Client, tableName string) *DynamoStore {
    return &DynamoStore{
        db: db,
        tableName: aws.String(tableName),
        now: func() time.Time {
            return time.Now().UTC()
        },
    }
}

