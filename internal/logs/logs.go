package logs

import (
	"context"
	"fmt"
)

var (
	defaultLogger Logger = &ConsoleLogger{}
)

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxFatal(ctx, fmt.Sprintf(fmt.Sprintf(format, v...)))
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxError(ctx, fmt.Sprintf(format, v...))
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxWarn(ctx, fmt.Sprintf(format, v...))
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxInfo(ctx, fmt.Sprintf(format, v...))
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxDebug(ctx, fmt.Sprintf(format, v...))
}

func SetLevel(ctx context.Context, level int) {
	defaultLogger.SetLevel(ctx, level)
}

func GetLevel(ctx context.Context) int {
	return defaultLogger.GetLevel(ctx)
}

func InitLogger(logger Logger) {
	defaultLogger = logger
}
