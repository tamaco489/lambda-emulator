package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	logger := logging.NewJSONLogger(ctx)

	logger.InfoContext(ctx, "processing EventBridge event",
		"source", event.Source,
		"detailType", event.DetailType,
		"account", event.AccountID,
		"region", event.Region,
	)

	var detail map[string]interface{}
	if err := json.Unmarshal(event.Detail, &detail); err != nil {
		logger.ErrorContext(ctx, "failed to parse event detail", "error", err)
		return err
	}

	logger.InfoContext(ctx, "event detail", "detail", detail)

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
