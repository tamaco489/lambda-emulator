package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

func TestNewLogger_ColdStart(t *testing.T) {
	coldStart = true

	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(context.Background(), handler)

	logger.Info("test message")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to parse log output: %v", err)
	}

	if got := result["coldStart"]; got != true {
		t.Errorf("coldStart = %v, want true", got)
	}
}

func TestNewLogger_WarmStart(t *testing.T) {
	coldStart = true
	_ = NewLogger(context.Background(), slog.NewJSONHandler(io.Discard, nil))

	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(context.Background(), handler)

	logger.Info("test message")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to parse log output: %v", err)
	}

	if got := result["coldStart"]; got != false {
		t.Errorf("coldStart = %v, want false", got)
	}
}

func TestNewLogger_WithLambdaContext(t *testing.T) {
	coldStart = true

	lc := &lambdacontext.LambdaContext{
		AwsRequestID:       "test-request-id",
		InvokedFunctionArn: "arn:aws:lambda:us-east-1:123456789012:function:test",
	}
	ctx := lambdacontext.NewContext(context.Background(), lc)

	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(ctx, handler)

	logger.Info("test message")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to parse log output: %v", err)
	}

	if got := result["requestId"]; got != "test-request-id" {
		t.Errorf("requestId = %v, want test-request-id", got)
	}
}

func TestNewJSONLogger(t *testing.T) {
	coldStart = true

	logger := NewJSONLogger(context.Background())
	if logger == nil {
		t.Fatal("NewJSONLogger() returned nil")
	}
}

func TestWithLogger_AndFromContext(t *testing.T) {
	coldStart = true

	logger := NewJSONLogger(context.Background())
	ctx := WithLogger(context.Background(), logger)

	retrieved := FromContext(ctx)
	if retrieved != logger {
		t.Error("FromContext() did not return the same logger")
	}
}

func TestFromContext_NoLogger(t *testing.T) {
	logger := FromContext(context.Background())
	if logger == nil {
		t.Fatal("FromContext() returned nil when no logger in context")
	}
}
