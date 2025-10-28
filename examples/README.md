# Examples

## サンプル一覧

| Example | 説明 | イベントタイプ |
|---------|------|--------------|
| basic | 最もシンプルな使用例 | カスタムイベント |
| alb | Application Load Balancerターゲットグループ | `events.ALBTargetGroupRequest` |
| apigateway | API Gatewayプロキシ統合 | `events.APIGatewayProxyRequest` |
| dynamodb | DynamoDB Streamsイベント処理 | `events.DynamoDBEvent` |
| eventbridge | EventBridgeカスタムイベント | `events.CloudWatchEvent` |
| kinesis | Kinesisデータストリーム処理 | `events.KinesisEvent` |
| s3 | S3 Putイベント処理 | `events.S3Event` |
| sqs | SQSキューからのメッセージ処理 | `events.SQSEvent` |

## 実行方法

**Terminal 1:**
```bash
cd examples/{example-name}
go run handler.go  # または main.go
```

**Terminal 2:**
```bash
go run cmd/lambda-emulator/main.go -event examples/{example-name}/event.json
```
