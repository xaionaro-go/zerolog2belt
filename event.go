package zerolog

import (
	"time"

	"github.com/facebookincubator/go-belt/tool/logger"
)

type Event struct {
	Level  logger.Level
	Logger *Logger
}

func (ev *Event) CallerSkipFrame(frameCount int) *Event {
	return ev
}

func (ev *Event) Msg(message string) {
	ev.Logger.Logger.Log(ev.Level, message)
}

func (ev *Event) Msgf(format string, args ...any) {
	ev.Logger.Logger.Logf(ev.Level, format, args...)
}

func (ev *Event) Str(key string, value string) *Event {
	return &Event{
		Level:  ev.Level,
		Logger: &Logger{Logger: ev.Logger.Logger.WithField(key, value)},
	}
}

func (ev *Event) Int(key string, value int) *Event {
	return &Event{
		Level:  ev.Level,
		Logger: &Logger{Logger: ev.Logger.Logger.WithField(key, value)},
	}
}

func (ev *Event) Dur(key string, value time.Duration) *Event {
	return &Event{
		Level:  ev.Level,
		Logger: &Logger{Logger: ev.Logger.Logger.WithField(key, value)},
	}
}

func (ev *Event) Err(err error) *Event {
	return &Event{Level: logger.LevelError, Logger: &Logger{ev.Logger.Logger.WithField("error", err)}}
}
