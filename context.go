package zerolog

import (
	"context"
)

type Context struct {
	logger *Logger
}

func Ctx(ctx context.Context) *Logger {
	return nil
}

func (ctx Context) Timestamp() Context {
	return ctx
}

func (ctx Context) Logger() *Logger {
	return ctx.logger
}
