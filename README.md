# lambda-emulator

> A tool to emulate AWS Lambda functions locally

### Table of Contents
- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Setup & Usage](#setup--usage)

### Overview

Solves the following challenges in local development with `aws-lambda-go`:

- **Missing RPC Client**: Need to write connection code to the RPC server started by `lambda.Start()` every time → Abstracted with `emulator`
- **Two-phase Protocol**: Need to understand and implement Ping→Invoke call sequence → Automatically executed internally
- **Insufficient Logging**: Lambda-specific information (cold start, request ID, etc.) not included in standard logs → Automatically added with `logging`
- **Deploy-first Approach**: Lightweight local testing difficult due to mandatory container/SAM → Instantly executable with CLI
- **No CLI Tool**: Need to write code for every test → Single command execution with `lambda-emulator`

### Prerequisites

- Go 1.25 or higher

### Setup

#### Installation
```bash
go install github.com/tamaco489/lambda-emulator/cmd/lambda-emulator@latest
```

#### Environment Variables
```bash
cp .env_sample .env
# Configure _LAMBDA_SERVER_PORT and other settings in .env
```

### Example Usage

Terminal 1: Start Lambda function
```bash
cd examples/dynamodb
go run handler.go
```

Terminal 2: Verify port 9000 is listening (optional)
```bash
lsof -i :9000
```

Terminal 2: Invoke with event
```bash
go run cmd/lambda-emulator/main.go -event examples/dynamodb/event.json
```

Terminal 1: Expected output
```json
{"time":"2025-10-29T02:00:16.415908544+09:00","level":"INFO","msg":"processing DynamoDB stream record","coldStart":true,"function":{"arn":""},"requestId":"","eventID":"1","eventName":"INSERT","tableName":"arn:aws:dynamodb:us-east-1:123456789012:table/MyTable/stream/2024-01-01T00:00:00.000"}
{"time":"2025-10-29T02:00:16.415952137+09:00","level":"INFO","msg":"new image","coldStart":true,"function":{"arn":""},"requestId":"","keys":{"Id":{"N":"101"}},"newImage":{"Age":{"N":"30"},"Id":{"N":"101"},"Name":{"S":"John Doe"}}}
```

#### Use as Library
```bash
go get github.com/tamaco489/lambda-emulator
```

See `examples/` directory for more examples (SQS, API Gateway, Kinesis, S3, EventBridge, etc.).
