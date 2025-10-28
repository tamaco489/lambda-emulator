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

type ResponseBody struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func handler(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	logger := logging.NewJSONLogger(ctx)
	logger.InfoContext(ctx, "received ALB request",
		"path", request.Path,
		"httpMethod", request.HTTPMethod,
	)

	respBody := ResponseBody{
		Message: "Hello from ALB",
		Path:    request.Path,
	}

	body, _ := json.Marshal(respBody)

	return events.ALBTargetGroupResponse{
		StatusCode:        200,
		StatusDescription: "200 OK",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(body),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
