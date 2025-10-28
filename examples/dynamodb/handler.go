package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	logger := logging.NewJSONLogger(ctx)

	for _, record := range event.Records {
		logger.InfoContext(ctx, "processing DynamoDB stream record",
			"eventID", record.EventID,
			"eventName", record.EventName,
			"tableName", record.EventSourceArn,
		)

		if record.Change.NewImage != nil {
			logger.InfoContext(ctx, "new image",
				"keys", record.Change.Keys,
				"newImage", record.Change.NewImage,
			)
		}

		if record.Change.OldImage != nil {
			logger.InfoContext(ctx, "old image",
				"oldImage", record.Change.OldImage,
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
