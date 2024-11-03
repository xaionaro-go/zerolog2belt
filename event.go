package zerolog

import (
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