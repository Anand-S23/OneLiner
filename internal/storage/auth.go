package storage

import (
	"context"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (store *DynamoStore) PutUser(user models.User) error {
    ur := models.NewUserRecord(user)
    item, err := attributevalue.MarshalMap(ur)
    if err != nil {
        return err
    }

    putItemInput := &dynamodb.PutItemInput {
        TableName: store.tableName,
        Item: item,
    }

    _, err = store.db.PutItem(context.TODO(), putItemInput)
    if err != nil {
        return err
    }

    return nil
}

