package logger

import (
	"flag"
	"testing"

	"github.com/polevpn/elog"
)

func TestLogger(t *testing.T) {
	flag.Parse()
	Info("sss", "xxxx")
	Infof("sss %v", "xxxx")
	elog.Flush()
}
