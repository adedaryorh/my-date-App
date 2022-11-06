package file

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type S3Client struct {
	client *s3.Client
	region string
}

func NewS3Client(config aws.Config, region string) *S3Client {
	return &S3Client{
		client: s3.NewFromConfig(config),
		region: region,
	}
}

func (sc *S3Client) UploadToBucket(ctx context.Context, bucket string, fileName string, file io.Reader) (*string, error) {
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	}

	_, err := sc.client.PutObject(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error uploading to bucket: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, sc.region, fileName)

	return &fileURL, nil
}
