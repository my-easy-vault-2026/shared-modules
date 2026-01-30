package logger

import "testing"

type TestLoggerTest struct {
	logger *testing.T
}

func NewTestLogger(t *testing.T) *TestLoggerTest {
	return &TestLoggerTest{logger: t}
}

func (t TestLoggerTest) Debug(args ...interface{}) {
	t.logger.Log(args...)
}

func (t TestLoggerTest) Debugf(format string, args ...interface{}) {
	t.logger.Logf(format, args...)
}

func (t TestLoggerTest) Info(args ...interface{}) {
	t.logger.Log(args...)
}

func (t TestLoggerTest) Infof(format string, args ...interface{}) {
	t.logger.Logf(format, args...)
}

func (t TestLoggerTest) Warn(args ...interface{}) {
	t.logger.Log(args...)
}

func (t TestLoggerTest) Warnf(format string, args ...interface{}) {
	t.logger.Logf(format, args...)
}

func (t TestLoggerTest) Error(args ...interface{}) {
	t.logger.Error(args...)
}

func (t TestLoggerTest) Errorf(format string, args ...interface{}) {
	t.logger.Errorf(format, args...)
}

func (t TestLoggerTest) Fatal(args ...interface{}) {
	t.logger.Fatal(args...)
}

func (t TestLoggerTest) Fatalf(format string, args ...interface{}) {
	t.logger.Fatalf(format, args...)
}

func (t TestLoggerTest) Println(args ...interface{}) {
	t.logger.Log(args...)
}

func (t TestLoggerTest) Printf(format string, args ...interface{}) {
	t.logger.Logf(format, args...)
}
