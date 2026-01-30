package logger

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

var LoggerInstance Logger

func Debug(args ...interface{}) {
	LoggerInstance.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	LoggerInstance.Debugf(format, args...)
}

func Info(args ...interface{}) {
	LoggerInstance.Info(args...)
}
func Infof(format string, args ...interface{}) {
	LoggerInstance.Infof(format, args...)
}

func Warn(args ...interface{}) {
	LoggerInstance.Warn(args...)
}
func Warnf(format string, args ...interface{}) {
	LoggerInstance.Warnf(format, args...)
}

func Error(args ...interface{}) {
	LoggerInstance.Error(args...)
}
func Errorf(format string, args ...interface{}) {
	LoggerInstance.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	LoggerInstance.Fatal(args...)
}
func Fatalf(format string, args ...interface{}) {
	LoggerInstance.Fatalf(format, args...)
}

func Println(args ...interface{}) {
	LoggerInstance.Println(args...)
}
func Printf(format string, args ...interface{}) {
	LoggerInstance.Printf(format, args...)
}
