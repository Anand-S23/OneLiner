package storage

import (
	"context"
	"log"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (store *SnippetStore) PutUser(user models.User) error {
    ur := models.NewUserRecord(user)
    item, err := attributevalue.MarshalMap(ur)
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

    log.Println("User successfully put into the db")
    return nil
}

func (store *SnippetStore) GetUser(pk string, sk string) models.User {
    input := &dynamodb.GetItemInput {
        TableName: store.tableName,
        ConsistentRead: aws.Bool(true),
        Key: map[string]types.AttributeValue{
            "PK":  &types.AttributeValueMemberS{Value: pk},
            "SK":  &types.AttributeValueMemberS{Value: sk},
        },
    }

    out, err := store.db.GetItem(context.TODO(), input)
    if err != nil {
        log.Printf("Could not get item for database, %s", err)
        return models.User{}
    }

    var ur models.UserRecord
    err = attributevalue.UnmarshalMap(out.Item, &ur)
    if err != nil {
        log.Printf("Could not unmarshal result item, %s", err)
        return models.User{}
    }

    user := models.NewUserFromRecord(ur)
    log.Printf("User returned from db: '%s'", user.ID)
    return user
}

