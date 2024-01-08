package storage

import (
	"context"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func idAndRng(id string, rng string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id":  &types.AttributeValueMemberS{Value: id},
		"rng": &types.AttributeValueMemberS{Value: rng},
	}
}

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

func (store *DynamoStore) GetUser(id string) (*models.User, error) {
    getItemInput := &dynamodb.GetItemInput {
        TableName: store.tableName,
        ConsistentRead: aws.Bool(true),
        Key: idAndRng(models.NewUserRecordHashKey(id), models.NewUserRecordRangeKey()),
    }

    out, err := store.db.GetItem(context.TODO(), getItemInput)
    if err != nil {
        return nil, err
    }

    var ur models.UserRecord
    err = attributevalue.UnmarshalMap(out.Item, &ur)
    if err != nil {
        return nil, err
    }

    user := models.NewUserFromRecord(ur)
    return user, nil
}

