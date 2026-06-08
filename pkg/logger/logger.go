package logger

import (
	"fmt"
	"log/slog"
)

type Logger struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Logger {
	return &Logger{log: log}
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Infof(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	l.log.Info(msg)
}

func (l *Logger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *Logger) Debugf(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	l.log.Debug(msg)
}

func (l *Logger) Error(msg string, err error) {
	l.log.Error(msg, slog.Any("err", err.Error()))
}

func (l *Logger) Errorf(msg string, err error, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	l.log.Error(msg, slog.Any("err", err.Error()))
}
