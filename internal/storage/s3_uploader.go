package storage

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage interface {
	Upload(ctx context.Context, key string, data []byte) error
}

type S3Uploader struct {
	client *s3.Client //AWS SDK client.
	bucket string
}

// constructor/ factory function
func NewS3Uploader(client *s3.Client, bucket string) *S3Uploader {
	return &S3Uploader{
		client: client,
		bucket: bucket,
	}
}

func (s *S3Uploader) Upload(ctx context.Context, key string, data []byte) error {

	_, err := s.client.PutObject(
		ctx, &s3.PutObjectInput{
			Bucket: &s.bucket,
			Key:    &key,
			Body:   bytes.NewReader(data), //Convert []byte into an io.Reader. The AWS SDK reads from this stream while uploading to S3.
		},
	)
	if err != nil {
		return err
	}

	return err
}
