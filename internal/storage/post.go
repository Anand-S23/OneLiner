package storage

import (
	"context"
	"io"
	"log"

	"github.com/Anand-S23/Snippet/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (store *SnippetStore) GetPost(postSK string) models.Post {
    keyCond := expression.KeyEqual(expression.Key("SK"), expression.Value(postSK))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		log.Printf("Failed to build query expression: %s", err)
        return models.Post{}
	}

	input := &dynamodb.QueryInput {
		TableName: store.tableName,
		IndexName: aws.String("GSI1"),
		KeyConditionExpression: expr.KeyCondition(),
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

    output, err := store.db.Query(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to get post by ID: %s", err)
        return models.Post{}
	}
	
	if len(output.Items) == 0 {
		log.Printf("Task not found")
        return models.Post{}
	}

    var pr models.PostRecord
    err = attributevalue.UnmarshalMap(output.Items[0], &pr)
    if err != nil {
        log.Printf("Could not unmarshal result item, %s", err)
        return models.Post{}
    }

    post := models.NewPostFromRecord(pr)
    log.Printf("Post returned from db: '%s'", post.ID)
    return post
}

func (store *SnippetStore) GetPostsByUser(postPK string) ([]models.Post, error) {
    input := &dynamodb.QueryInput {
        TableName: store.tableName,
        ConsistentRead: aws.Bool(true),
        KeyConditionExpression: aws.String("PK = :pk"),
        ExpressionAttributeValues: map[string]types.AttributeValue {
            ":pk": &types.AttributeValueMemberS{Value: postPK},
        },
    }

    output, err := store.db.Query(context.TODO(), input)
	if err != nil {
		return []models.Post{}, err
	}

    var posts []models.Post
    for _, item := range output.Items {
        var record models.Record

        err = attributevalue.UnmarshalMap(item, &record)
        if err != nil {
            return []models.Post{}, err
        }

        switch record.Type {
        case models.PostRecordType:
            var post models.Post
            err = attributevalue.UnmarshalMap(item, &post)
			if err != nil {
                return []models.Post{}, err
			}

            posts = append(posts, post)
        }
    }

    return posts, nil
}

func (store *SnippetStore) DeletePost(pk string, sk string) error {
    _, err := store.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
        TableName: store.tableName,
        Key: map[string]types.AttributeValue{
            "PK":  &types.AttributeValueMemberS{Value: pk},
            "SK":  &types.AttributeValueMemberS{Value: sk},
        },
    })

    return err
}

