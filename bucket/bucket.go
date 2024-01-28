package bucket

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListURLs() []string {
	bucket := os.Getenv("AWS_BUCKET_NAME")
	region := "us-east-1"

	sess, err := session.NewSession(
		&aws.Config{
			Region: &region,
		},
	)

	if err != nil {
		panic(err)
	}

	s3Bucket := s3.New(sess)

	resp, err := s3Bucket.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: &bucket,
			Prefix: aws.String(os.Getenv("AWS_BUCKET_PREFIX")),
		},
	)

	if err != nil {
		panic(err)
	}

	images := make([]string, 0, len(resp.Contents))

	for i, r := range resp.Contents {
		if i == 0 {
			continue
		}

		images = append(images, fmt.Sprintf("%s/%s", os.Getenv("AWS_BUCKET_URL"), *r.Key))
	}

	return images
}
