package zerolog

import (
	"fmt"
	"io"

	"github.com/facebookincubator/go-belt/tool/logger"
	"github.com/facebookincubator/go-belt/tool/logger/adapter"
)

type Logger struct {
	logger.Logger
}

func New(w io.Writer) *Logger {
	return &Logger{
		Logger: adapter.LoggerFromPrintfer(printfer{w: w}),
	}
}

type printfer struct {
	w io.Writer
}

func (p printfer) Printf(format string, args ...any) {
	fmt.Fprintf(p.w, format, args...)
}

func (l *Logger) With() Context {
	return Context{logger: l}
}

func (l *Logger) Output(w io.Writer) Logger {
	return *New(w)
}

func (l *Logger) Level(level Level) Logger {
	return Logger{l.Logger.WithLevel(level)}
}

func (l *Logger) Sample(sampler Sampler) Logger {
	return Logger{l.Logger.WithHooks(sampler)}
}

func (l *Logger) Hook(hook Hook) Logger {
	return Logger{l.Logger.WithHooks(hook)}
}

func (l *Logger) Err(err error) *Event {
	return &Event{Level: logger.LevelError, Logger: &Logger{l.Logger.WithField("error", err)}}
}

func (l *Logger) Trace() *Event {
	return &Event{Level: logger.LevelTrace, Logger: l}
}

func (l *Logger) Debug() *Event {
	return &Event{Level: logger.LevelDebug, Logger: l}
}

func (l *Logger) Info() *Event {
	return &Event{Level: logger.LevelInfo, Logger: l}
}

func (l *Logger) Warn() *Event {
	return &Event{Level: logger.LevelWarning, Logger: l}
}

func (l *Logger) Error() *Event {
	return &Event{Level: logger.LevelError, Logger: l}
}

func (l *Logger) Panic() *Event {
	return &Event{Level: logger.LevelPanic, Logger: l}
}

func (l *Logger) Fatal() *Event {
	return &Event{Level: logger.LevelFatal, Logger: l}
}

func (l *Logger) WithLevel(level logger.Level) *Event {
	return &Event{Level: level, Logger: l}
}

func (l *Logger) Log() *Event {
	return &Event{Logger: l}
}
