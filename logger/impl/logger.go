package impl

import (
	"flag"
	"fmt"
	"regexp"
	"runtime"

	"shared-modules/logger"

	// "github.com/petermattis/goid"
	"github.com/polevpn/elog"
	"gitlab.com/jjjj.sinrui2023/goid"

	"shared-modules/utils"
)

type LoggerImpl struct {
	lineFolding   bool
	re            *regexp.Regexp
	replaceString string
}

var _ logger.Logger = (*LoggerImpl)(nil)

func init() {
	instance := &LoggerImpl{}
	flag.BoolVar(&instance.lineFolding, "lineFolding", false, "fold to one line,default false")
	if instance.lineFolding {
		instance.replaceString = " "
	} else {
		instance.replaceString = "\r\n"
	}
	instance.re = regexp.MustCompile(`\r?\n`)
	logger.LoggerInstance = instance
}

func (l LoggerImpl) Debug(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Debug(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Debugf(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Debug(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}

func (l LoggerImpl) Info(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Info(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Infof(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Info(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}

func (l LoggerImpl) Warn(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	for i, a := range args {
		if _, ok := a.(error); ok {
			args[i] = utils.FormatErrorChainLink(a.(error))
		}
	}
	elog.Warn(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Warnf(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	for i, a := range args {
		if _, ok := a.(error); ok {
			args[i] = utils.FormatErrorChainLink(a.(error))
		}
	}
	elog.Warn(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}

func (l LoggerImpl) Error(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	for i, a := range args {
		if _, ok := a.(error); ok {
			args[i] = utils.FormatErrorChainLink(a.(error))
		}
	}
	elog.Error(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Errorf(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	for i, a := range args {
		if _, ok := a.(error); ok {
			args[i] = utils.FormatErrorChainLink(a.(error))
		}
	}
	elog.Error(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}

func (l LoggerImpl) Fatal(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Fatal(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Fatalf(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Fatal(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}

func (l LoggerImpl) Println(args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Println(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprint(args...), l.replaceString), filepath, line))
}
func (l LoggerImpl) Printf(format string, args ...interface{}) {
	_, filepath, line, _ := runtime.Caller(2)
	reqId, _ := utils.GetMDCValue("reqId")
	elog.Println(fmt.Sprintf("\x1b[37;44m%12s\x1b[0;0m \x1b[35;100m%d\x1b[0;0m %s \x1b[38;5;245m(%s:%d)\x1b[0m", reqId, goid.Get(), l.re.ReplaceAllString(fmt.Sprintf(format, args...), l.replaceString), filepath, line))
}
