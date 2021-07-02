package logrus_formatter

import "github.com/petermattis/goid"

func getGoId() int64 {
	return goid.Get()
}
