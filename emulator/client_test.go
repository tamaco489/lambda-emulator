package emulator

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func TestClient_Invoke(t *testing.T) {
	port := 19000
	go startTestLambda(t, port)
	time.Sleep(100 * time.Millisecond)

	client, err := NewClient(&Config{Port: port})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	tests := []struct {
		name    string
		payload []byte
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			payload: []byte(`{"message":"hello"}`),
			want:    `"hello"`,
			wantErr: false,
		},
		{
			name:    "empty payload",
			payload: []byte(`{}`),
			want:    `""`,
			wantErr: false,
		},
		{
			name:    "complex payload",
			payload: []byte(`{"message":"world","count":42}`),
			want:    `"world"`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := client.Invoke(ctx, tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("Invoke() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestClientWithOptions(t *testing.T) {
	port := 19001
	go startTestLambda(t, port)
	time.Sleep(100 * time.Millisecond)

	client, err := NewClientWithOptions(
		WithPort(port),
		WithInvokeTimeout(5*time.Minute),
		WithConnectTimeout(5*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to create client with options: %v", err)
	}
	defer client.Close()

	payload := []byte(`{"message":"test"}`)
	result, err := client.Invoke(context.Background(), payload)
	if err != nil {
		t.Fatalf("Invoke() failed: %v", err)
	}

	expected := `"test"`
	if string(result) != expected {
		t.Errorf("Invoke() = %v, want %v", string(result), expected)
	}
}

func startTestLambda(t *testing.T, port int) {
	t.Helper()

	os.Setenv("_LAMBDA_SERVER_PORT", fmt.Sprintf("%d", port))

	handler := func(ctx context.Context, event map[string]interface{}) (string, error) {
		if msg, ok := event["message"].(string); ok {
			return msg, nil
		}
		return "", nil
	}

	lambda.Start(handler)
}
