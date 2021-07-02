package json_formatter

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			"FieldKeyTime":  "@timestamp",
			"FieldKeyLevel": "@level",
			"FieldKeyMsg":   "@message",
			"FieldKeyFunc":  "@caller",
		},
		//DataKey: "hhhhhh",
	})

	logrus.SetOutput(os.Stdout)

	logrus.WithField("ctxId", "ddddd").Info("Sdfadsa")

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")

}
