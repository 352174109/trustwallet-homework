package logs

import (
	"context"
	"log"
	"sync/atomic"
)

type Logger interface {
	CtxDebug(ctx context.Context, msg string)

	CtxInfo(ctx context.Context, msg string)

	CtxWarn(ctx context.Context, msg string)

	CtxError(ctx context.Context, msg string)

	CtxFatal(ctx context.Context, msg string)

	SetLevel(ctx context.Context, level int)

	GetLevel(ctx context.Context) int
}

type ConsoleLogger struct {
	level int32
}

func (c *ConsoleLogger) CtxDebug(ctx context.Context, msg string) {
	if LevelDebug < c.GetLevel(ctx) {
		return
	}

	log.Println("DEBUG: " + msg)
}

func (c *ConsoleLogger) CtxInfo(ctx context.Context, msg string) {
	if LevelInfo < c.GetLevel(ctx) {
		return
	}

	log.Println("INFO: " + msg)
}

func (c *ConsoleLogger) CtxWarn(ctx context.Context, msg string) {
	if LevelWarn < c.GetLevel(ctx) {
		return
	}
	log.Println("WARN: " + msg)
}

func (c *ConsoleLogger) CtxError(ctx context.Context, msg string) {
	if LevelError < c.GetLevel(ctx) {
		return
	}
	log.Println("ERROR: " + msg)
}

func (c *ConsoleLogger) CtxFatal(ctx context.Context, msg string) {
	if LevelFatal < c.GetLevel(ctx) {
		return
	}
	log.Fatalln(msg)
}

func (c *ConsoleLogger) SetLevel(ctx context.Context, level int) {
	atomic.StoreInt32(&c.level, int32(level))
}

func (c *ConsoleLogger) GetLevel(ctx context.Context) int {
	return int(atomic.LoadInt32(&c.level))
}
