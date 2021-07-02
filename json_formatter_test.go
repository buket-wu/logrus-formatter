package logrus_formatter

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	f := NewJsonFormatter(false)

	logrus.SetFormatter(f)

	logrus.SetOutput(os.Stdout)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")

}
