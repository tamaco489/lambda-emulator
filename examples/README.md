# Examples

## Example List

| Example | Description | Event Type |
|---------|-------------|------------|
| basic | Simplest usage example | Custom event |
| alb | Application Load Balancer target group | `events.ALBTargetGroupRequest` |
| apigateway | API Gateway proxy integration | `events.APIGatewayProxyRequest` |
| dynamodb | DynamoDB Streams event processing | `events.DynamoDBEvent` |
| eventbridge | EventBridge custom event | `events.CloudWatchEvent` |
| kinesis | Kinesis Data Streams processing | `events.KinesisEvent` |
| s3 | S3 Put event processing | `events.S3Event` |
| sqs | SQS queue message processing | `events.SQSEvent` |

## How to Run

**Terminal 1:**
```bash
cd examples/{example-name}
go run handler.go  # or main.go
```

**Terminal 2:**
```bash
go run cmd/lambda-emulator/main.go -event examples/{example-name}/event.json
```
