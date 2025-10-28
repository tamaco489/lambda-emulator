package emulator

import "time"

type Option func(*Config)

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithInvokeTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.InvokeTimeout = timeout
	}
}

func WithConnectTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.ConnectTimeout = timeout
	}
}

func NewClientWithOptions(opts ...Option) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return NewClient(cfg)
}
