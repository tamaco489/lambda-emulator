package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/tamaco489/lambda-emulator/logging"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Reply string `json:"reply"`
}

func handler(ctx context.Context, req Request) (Response, error) {
	logger := logging.NewJSONLogger(ctx)
	logger.InfoContext(ctx, "received request", "message", req.Message)

	return Response{
		Reply: "Hello, " + req.Message,
	}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
