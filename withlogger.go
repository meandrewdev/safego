package safego

import (
	"runtime/debug"

	"github.com/meandrewdev/logger"
)

func WithDefaultLogger(f func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if panicMessage := recover(); panicMessage != nil {
				stack := debug.Stack()

				logger.MessageF("PANIC: %v\nSTACK: %s", logger.LG_Error, panicMessage, stack)
			}
		}()

		f(args...)
	}()
}

func WithLogger(l *logger.Logger, f func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if panicMessage := recover(); panicMessage != nil {
				stack := debug.Stack()

				l.MessageF("PANIC: %v\nSTACK: %s", logger.LG_Error, panicMessage, stack)
			}
		}()

		f(args...)
	}()
}
