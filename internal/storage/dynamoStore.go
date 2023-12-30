package storage

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoStore struct {
	db *dynamodb.Client
}

func NewDynamoStore(db *dynamodb.Client) *DynamoStore {
    return &DynamoStore{db: db}
}

