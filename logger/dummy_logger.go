package logger

import (
	"fmt"
)

type DummyLogger struct {
}

var _ Logger = (*DummyLogger)(nil)

func (l DummyLogger) Debug(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l DummyLogger) Info(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l DummyLogger) Warn(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l DummyLogger) Error(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l DummyLogger) Fatal(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}

func (l DummyLogger) Println(args ...interface{}) {
	fmt.Println(args)
}
func (l DummyLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args)
}
