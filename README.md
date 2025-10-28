#  lambda-emulator

> AWS Lambda関数をローカルでエミュレートするツール

## 目次
- [概要](#概要)
- [前提条件](#前提条件)
- [セットアップ・使い方](#セットアップ・使い方)

## 概要

`aws-lambda-go`を使ったローカル開発における以下の課題を解決する

- **RPCクライアントの不在**: `lambda.Start()`で起動するRPCサーバーへの接続コードを毎回書く必要がある → `emulator`で抽象化
- **2段階プロトコル**: Ping→Invokeの呼び出し順序を理解・実装する必要がある → 内部で自動実行
- **ログの不足**: Lambda固有情報（コールドスタート、リクエストID等）が標準ログに含まれない → `logging`で自動付与
- **デプロイ中心**: コンテナ/SAM必須で軽量なローカルテストが困難 → CLIで即座に実行可能
- **CLIツール不在**: テストのたびにコードを書く必要がある → `lambda-emulator`コマンド一発実行

## 前提条件

- Go 1.25 以上

## セットアップ・使い方

### インストール
```bash
go install github.com/tamaco489/lambda-emulator/cmd/lambda-emulator@latest
```

### 環境変数設定
```bash
cp .env_sample .env
# .envで_LAMBDA_SERVER_PORTなどを設定
```

### 使い方
```bash
# Terminal 1: Lambda関数起動（.envから環境変数読み込み）
go run main.go

# Terminal 2: 実行
lambda-emulator -event event.json
```

### ライブラリとして使用
```bash
go get github.com/tamaco489/lambda-emulator
```

詳細は `examples/` ディレクトリを参照。
