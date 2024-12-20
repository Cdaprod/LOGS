// logging/zap_logger.go

package zaplogger

import (
	"context"
	"go.uber.org/zap"
	"yourproject/logging"
)

// ZapLogger implements the Logger interface using Uber's zap logging library.
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger creates a new ZapLogger instance.
func NewZapLogger() (*ZapLogger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{logger: zapLogger}, nil
}

// Info logs an informational message.
func (z *ZapLogger) Info(ctx context.Context, message string, fields ...map[string]interface{}) {
	z.logger.Info(message, zap.Any("fields", fields))
}

// Implement other methods (Warn, Error, Debug) similarly...

// Close the logger, flush any buffered log entries.
func (z *ZapLogger) Close() error {
	return z.logger.Sync()
}