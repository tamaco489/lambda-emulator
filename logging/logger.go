package logging

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

var (
	coldStart bool = true
	mu        sync.Mutex
)

func NewLogger(ctx context.Context, handler slog.Handler) *slog.Logger {
	attrs := extractLambdaAttributes(ctx)
	return slog.New(handler).With(attrs...)
}

func NewJSONLogger(ctx context.Context) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return NewLogger(ctx, handler)
}

func extractLambdaAttributes(ctx context.Context) []any {
	attrs := make([]any, 0)

	mu.Lock()
	isColdStart := coldStart
	if coldStart {
		coldStart = false
	}
	mu.Unlock()
	attrs = append(attrs, slog.Bool("coldStart", isColdStart))

	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		attrs = append(attrs,
			slog.Group("function",
				slog.String("arn", lc.InvokedFunctionArn),
			),
			slog.String("requestId", lc.AwsRequestID),
		)
		return attrs
	}

	name := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	if name == "" {
		return attrs
	}

	attrs = append(attrs,
		slog.Group("function",
			slog.String("name", name),
			slog.String("version", os.Getenv("AWS_LAMBDA_FUNCTION_VERSION")),
		),
	)

	return attrs
}
