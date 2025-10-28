package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

func handler(ctx context.Context, event events.S3Event) error {
	logger := logging.NewJSONLogger(ctx)

	for _, record := range event.Records {
		logger.InfoContext(ctx, "processing S3 event",
			"eventName", record.EventName,
			"bucket", record.S3.Bucket.Name,
			"key", record.S3.Object.Key,
			"size", record.S3.Object.Size,
			"etag", record.S3.Object.ETag,
		)

		if record.EventName == "ObjectCreated:Put" {
			logger.InfoContext(ctx, "new object uploaded",
				"bucket", record.S3.Bucket.Name,
				"key", record.S3.Object.Key,
			)
		}
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
