package mongodb

import (
	"testing"
)

type ucTest struct {
	in, out string
}

var fss FolderStatService

func TestDailyFlushing(t *testing.T) {
	fss.DailyFlushing()
}
