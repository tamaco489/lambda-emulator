package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

func handler(ctx context.Context, event events.SQSEvent) error {
	logger := logging.NewJSONLogger(ctx)

	for _, record := range event.Records {
		logger.InfoContext(ctx, "processing SQS message",
			"messageId", record.MessageId,
			"body", record.Body,
			"eventSource", record.EventSource,
		)

		if err := processMessage(record); err != nil {
			logger.ErrorContext(ctx, "failed to process message",
				"error", err,
				"messageId", record.MessageId,
			)
			return err
		}
	}

	return nil
}

func processMessage(record events.SQSMessage) error {
	fmt.Printf("Processing: %s\n", record.Body)
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
