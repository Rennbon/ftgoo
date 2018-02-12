package tool

import (
	"time"
)

var maxUnix int64 = time.Date(3000, 1, 1, 0, 0, 0, 0, time.Local).Unix()
var minUnix = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local).Unix()

func GetDate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func CheckUnix(unix int64) bool {
	if unix > maxUnix || unix < minUnix {
		return false
	} else {
		return true
	}
}
