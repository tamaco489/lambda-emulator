package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

func handler(ctx context.Context, event events.KinesisEvent) error {
	logger := logging.NewJSONLogger(ctx)

	for _, record := range event.Records {
		logger.InfoContext(ctx, "processing Kinesis record",
			"eventID", record.EventID,
			"eventName", record.EventName,
			"sequenceNumber", record.Kinesis.SequenceNumber,
			"partitionKey", record.Kinesis.PartitionKey,
			"data", string(record.Kinesis.Data),
		)
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
