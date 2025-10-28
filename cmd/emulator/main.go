package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tamaco489/lambda-emulator/emulator"
)

const version = "0.1.0"

var (
	portFlag    = flag.Int("port", 9000, "Lambda RPC port")
	eventFlag   = flag.String("event", "", "Event JSON file path (required)")
	timeoutFlag = flag.Duration("timeout", 15*time.Minute, "Invocation timeout")
	versionFlag = flag.Bool("version", false, "Print version")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf("lambda-emulator v%s\n", version)
		os.Exit(0)
	}

	if *eventFlag == "" {
		log.Fatal("Error: -event flag is required\n\nUsage: lambda-emulator -event <file>")
	}

	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	payload, err := os.ReadFile(*eventFlag)
	if err != nil {
		return fmt.Errorf("failed to read event file: %w", err)
	}

	client, err := emulator.NewClientWithOptions(
		emulator.WithPort(*portFlag),
		emulator.WithInvokeTimeout(*timeoutFlag),
	)
	if err != nil {
		return fmt.Errorf("failed to create emulator client: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), *timeoutFlag)
	defer cancel()

	result, err := client.Invoke(ctx, payload)
	if err != nil {
		return fmt.Errorf("invocation failed: %w", err)
	}

	fmt.Println(string(result))
	return nil
}
