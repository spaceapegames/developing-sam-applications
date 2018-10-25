package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var destinationBucket = os.Getenv("DESTINATION_BUCKET")

func Replicate(event events.S3Event) error {
	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.Key

	log.Printf("replicating %s from bucket %s, to %s", key, bucket, destinationBucket)

	svc := s3.New(session.Must(session.NewSession()))
	_, err := svc.CopyObject(
		&s3.CopyObjectInput{
			Bucket: aws.String(destinationBucket),
			CopySource: aws.String(bucket + "/" + key),
			Key: aws.String(key),
	})

	return err
}

func main() {
	lambda.Start(Replicate)
}