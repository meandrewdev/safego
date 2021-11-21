package safego

import (
	"fmt"

	"github.com/meandrewdev/logger"
	"github.com/ztrue/tracerr"
)

func WithDefaultLogger(f func(args ...interface{}), args ...interface{}) {
	WithLogger(logger.GetInstance(), f, args...)
}

func WithLogger(l logger.Logger, f func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if panicMessage := recover(); panicMessage != nil {
				err := tracerr.New(fmt.Sprintf("PANIC: %v", panicMessage))
				logger.Error(err)
			}
		}()

		f(args...)
	}()
}
