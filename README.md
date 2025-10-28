# lambda-emulator

> A tool to emulate AWS Lambda functions locally

## Table of Contents
- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Setup & Usage](#setup--usage)

## Overview

Solves the following challenges in local development with `aws-lambda-go`:

- **Missing RPC Client**: Need to write connection code to the RPC server started by `lambda.Start()` every time → Abstracted with `emulator`
- **Two-phase Protocol**: Need to understand and implement Ping→Invoke call sequence → Automatically executed internally
- **Insufficient Logging**: Lambda-specific information (cold start, request ID, etc.) not included in standard logs → Automatically added with `logging`
- **Deploy-first Approach**: Lightweight local testing difficult due to mandatory container/SAM → Instantly executable with CLI
- **No CLI Tool**: Need to write code for every test → Single command execution with `lambda-emulator`

## Prerequisites

- Go 1.25 or higher

## Setup & Usage

### Installation
```bash
go install github.com/tamaco489/lambda-emulator/cmd/lambda-emulator@latest
```

### Environment Variables
```bash
cp .env_sample .env
# Configure _LAMBDA_SERVER_PORT and other settings in .env
```

### Usage
```bash
# Terminal 1: Start Lambda function (loads environment variables from .env)
go run main.go

# Terminal 2: Execute
lambda-emulator -event event.json
```

### Use as Library
```bash
go get github.com/tamaco489/lambda-emulator
```

See `examples/` directory for details.
