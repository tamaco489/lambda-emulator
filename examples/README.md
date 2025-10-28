# Examples

## セットアップ

各exampleディレクトリで.envファイルを作成：

```bash
cd examples/basic  # または examples/sqs
cp ../../.env_sample .env
```

## Basic Example

最もシンプルな使用例。

### 実行方法

**Terminal 1:**
```bash
cd examples/basic
go run main.go
```

**Terminal 2:**
```bash
go run cmd/lambda-emulator/main.go -event examples/basic/event.json
```

**期待される出力:**
```json
{"reply":"Hello, World"}
```

---

## SQS Example

SQSイベントを処理する例。

### 実行方法

**Terminal 1:**
```bash
cd examples/sqs
go run handler.go
```

**Terminal 2:**
```bash
go run cmd/lambda-emulator/main.go -event examples/sqs/event.json
```

**期待される出力:**
- Terminal 1にSQSメッセージ処理ログが表示される
- Terminal 2に空のレスポンスが返る
