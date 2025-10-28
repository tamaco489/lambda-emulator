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

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := logging.NewJSONLogger(ctx)
	logger.InfoContext(ctx, "received API Gateway request",
		"path", request.Path,
		"httpMethod", request.HTTPMethod,
	)

	var reqBody RequestBody
	if err := json.Unmarshal([]byte(request.Body), &reqBody); err != nil {
		logger.ErrorContext(ctx, "failed to parse request body", "error", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error":"Invalid request body"}`,
		}, nil
	}

	respBody := ResponseBody{
		Message: "Hello, " + reqBody.Name,
	}

	body, _ := json.Marshal(respBody)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	lambda.Start(handler)
}
