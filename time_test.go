package goutil

import (
	"testing"
	"time"
)

func TestGetCurrentTimeStr(t *testing.T) {
	expectStr := time.Now().Format("2006-01-02 15:04:05")
	actualStr := GetCurrentTimeStr()

	if expectStr != actualStr {
		t.Errorf("GetCurrentTimeStr() func error,actual:%s expect:%s", actualStr, expectStr)
	}
}
