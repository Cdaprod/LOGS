// main.go

package main

import (
	"context"
	"log"
	"yourproject/logging"
	"yourproject/logging/zaplogger"
)

func main() {
	// Initialize the logger (could be based on configuration settings)
	logger, err := zaplogger.NewZapLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Example usage of logger in a service
	ctx := context.Background()
	logger.Info(ctx, "Starting the registry service", map[string]interface{}{"service": "REGS"})

	// ... rest of the application
}