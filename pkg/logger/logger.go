package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type LoggerKeyType string

var logKey LoggerKeyType = "Logger"

// To store and retrieve the logger details
func Inject(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, log)
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(logKey).(*zap.Logger)

	if !ok {
		fmt.Println("Logger from Context Error ")
	}

	return log
}
