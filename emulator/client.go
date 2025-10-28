package emulator

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

type Client struct {
	rpcClient *rpc.Client
	port      int
}

type Config struct {
	Port           int
	InvokeTimeout  time.Duration
	ConnectTimeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Port:           9000,
		InvokeTimeout:  15 * time.Minute,
		ConnectTimeout: 10 * time.Second,
	}
}

func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	conn, err := net.DialTimeout("tcp",
		fmt.Sprintf("localhost:%d", cfg.Port),
		cfg.ConnectTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Lambda runtime: %w", err)
	}

	rpcClient := rpc.NewClient(conn)
	return &Client{
		rpcClient: rpcClient,
		port:      cfg.Port,
	}, nil
}

func (c *Client) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	if err := c.ping(); err != nil {
		return nil, fmt.Errorf("health check failed: %w", err)
	}

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(15 * time.Minute)
	}

	req := messages.InvokeRequest{
		Payload: payload,
		Deadline: messages.InvokeRequest_Timestamp{
			Seconds: deadline.Unix(),
			Nanos:   int64(deadline.Nanosecond()),
		},
	}

	var res messages.InvokeResponse
	if err := c.rpcClient.Call("Function.Invoke", req, &res); err != nil {
		return nil, fmt.Errorf("invocation failed: %w", err)
	}

	if res.Error != nil {
		return nil, fmt.Errorf("function error: %s", res.Error.Message)
	}

	return res.Payload, nil
}

func (c *Client) Close() error {
	if c.rpcClient != nil {
		return c.rpcClient.Close()
	}
	return nil
}

func (c *Client) ping() error {
	var req messages.PingRequest
	var res messages.PingResponse
	return c.rpcClient.Call("Function.Ping", &req, &res)
}
