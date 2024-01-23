package blob

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Bucket struct {
    BucketName *string
    Client     *s3.Client
    Uploader   *manager.Uploader
    Downloader *manager.Downloader
}

func InitBlob(bucketName string, timeout time.Duration) *S3Bucket {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    cfg, err := config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
        opts.Region = "us-east-1"
        return nil
    })

    client := s3.NewFromConfig(cfg)
    if err != nil {
        log.Panic(err);
    }

    createS3Bucket(client, bucketName, timeout)

    return &S3Bucket {
        BucketName: aws.String(bucketName),
        Client: client,
        Uploader: manager.NewUploader(client),
        Downloader: manager.NewDownloader(client),
    }
}

func createS3Bucket(client *s3.Client, bucketName string, timeout time.Duration) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

	buckets, err := client.ListBuckets(ctx, nil)
	if err != nil {
        log.Fatalf("Could not create bucket: %s\n", err)
	}

    if (bucketExists(buckets, bucketName)) {
        log.Println("Skipping creation of bucket: already exists")
    } else {
        _, err = client.CreateBucket(ctx, &s3.CreateBucketInput {
            Bucket: aws.String(bucketName),
        })
        if err != nil {
            log.Fatalf("Could not create bucket: %s\n", err)
        }
        log.Println("Bucket created sucessfully")
    }
}

func bucketExists(buckets *s3.ListBucketsOutput, bucketName string) bool {
    for _, b := range buckets.Buckets {
        if *b.Name == bucketName {
            return true
        }
    }

    return false
}

