package database

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var usersTableSchema *dynamodb.CreateTableInput = &dynamodb.CreateTableInput {
		TableName: aws.String("Users"),
		AttributeDefinitions: []types.AttributeDefinition {
			{
				AttributeName: aws.String("username"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("password"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("pfp"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement {
			{
				AttributeName: aws.String("username"),
				KeyType: types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("email"),
				KeyType: types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput {
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
}

var postsTableSchema *dynamodb.CreateTableInput = &dynamodb.CreateTableInput {
		TableName: aws.String("Posts"),
		AttributeDefinitions: []types.AttributeDefinition {
            {
                AttributeName: aws.String("id"),
                AttributeType: types.ScalarAttributeTypeS,
            },
			{
				AttributeName: aws.String("body"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("user"),
				AttributeType: types.ScalarAttributeTypeS,
			},
            {
                AttributeName: aws.String("type"),
				AttributeType: types.ScalarAttributeTypeS,
            },
            {
                AttributeName: aws.String("refrence"),
                AttributeType: types.ScalarAttributeTypeS,
            },
            {
                AttributeName: aws.String("likes"),
                AttributeType: types.ScalarAttributeTypeN,
            },
            {
                AttributeName: aws.String("uses"),
                AttributeType: types.ScalarAttributeTypeN,
            },
		},
		KeySchema: []types.KeySchemaElement {
			{
				AttributeName: aws.String("id"),
				KeyType: types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput {
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
}

func InitDB(timeout time.Duration) *dynamodb.Client {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    cfg, err := config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
        opts.Region = "us-east-1"
        return nil
    })

    db := dynamodb.NewFromConfig(cfg)
    if err != nil {
        log.Panic(err);
    }

    createDynamoDBTable(db, usersTableSchema, timeout)
    createDynamoDBTable(db, postsTableSchema, timeout)
    return db
}

func createDynamoDBTable(db *dynamodb.Client, input *dynamodb.CreateTableInput, timeout time.Duration) {
	var tableDesc *types.TableDescription
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

	table, err := db.CreateTable(ctx, input)
	if err != nil {
		log.Fatalf("Failed to create table %v with error: %v\n", input.TableName, err)
	} 

    waiter := dynamodb.NewTableExistsWaiter(db)

    err = waiter.Wait(ctx, &dynamodb.DescribeTableInput { 
        TableName: aws.String(*input.TableName),
    }, 5 * time.Minute)
    if err != nil {
        log.Printf("Failed to wait on create table %v with error: %v\n", input.TableName, err)
    }

    tableDesc = table.TableDescription
    log.Printf("%v table created sucessfully: %v\n", input.TableName, tableDesc)
}

