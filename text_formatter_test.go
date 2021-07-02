package logrus_formatter

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestFormatter_print(t *testing.T) {
	//SetCtxId(uuid.New())
	f := NewFormatter(false)
	//f.SetCtxId(uuid.NewString())

	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.SetFormatter(f)

	logrus.SetOutput(os.Stdout)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
}
