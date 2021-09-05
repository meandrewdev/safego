package safego

import "runtime/debug"

func WithFunc(handler func(message interface{}, stack []byte), f func(args ...interface{}), args ...interface{}) {
	go func() {
		defer func() {
			if panicMessage := recover(); panicMessage != nil {
				stack := debug.Stack()
				handler(panicMessage, stack)
			}
		}()

		f(args...)
	}()
}
