package logger

import (
	"context"

	"go.uber.org/zap"
)

type LoggerKeyType string

var logKey LoggerKeyType = "Logger"

func Inject(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, log)
}
