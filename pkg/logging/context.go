package logging

import (
	"context"
	"os"
)

type ctxKey struct{}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(ctxKey{}).(*Logger)
	if !ok {
		return NewLogger(os.Stderr)
	}
	return logger
}
