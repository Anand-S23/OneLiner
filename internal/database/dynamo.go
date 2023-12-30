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

var UsersTableName string = "OLUsers"
var PostsTableName string = "OLPosts"

var usersTableSchema *dynamodb.CreateTableInput = &dynamodb.CreateTableInput {
		TableName: aws.String(UsersTableName),
		AttributeDefinitions: []types.AttributeDefinition {
			{
				AttributeName: aws.String("username"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("email"),
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
		TableName: aws.String(PostsTableName),
		AttributeDefinitions: []types.AttributeDefinition {
            {
                AttributeName: aws.String("id"),
                AttributeType: types.ScalarAttributeTypeS,
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

    result, err := db.ListTables(ctx, &dynamodb.ListTablesInput{})
    if err != nil {
        log.Fatalf("Could not get tables: %s", err)
    }
    tables := result.TableNames
    
    if !tableExists(tables, *usersTableSchema.TableName) {
        err = createDynamoDBTable(db, usersTableSchema, timeout)
        if err != nil {
            log.Fatalf("Could not create Users table: %s\n", err)
        }
        log.Println("Users table created sucessfully")
    } else {
        log.Println("Skipping creation of Users table: already exists")
    }

    if !tableExists(tables, *postsTableSchema.TableName) {
        createDynamoDBTable(db, postsTableSchema, timeout)
        if err != nil {
            log.Fatalf("Could not create Posts table: %s\n", err)
        }
        log.Println("Posts table created sucessfully")
    } else {
        log.Println("Skipping creation of Posts table: already exists")
    }

    return db
}

func createDynamoDBTable(db *dynamodb.Client, input *dynamodb.CreateTableInput, timeout time.Duration) error {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

	_, err := db.CreateTable(ctx, input)
	if err != nil {
        return err
	} 

    return nil
}

func tableExists(tables []string, tableName string) bool {
    for _, t := range tables {
        if t == tableName {
            return true
        }
    }

    return false
}

