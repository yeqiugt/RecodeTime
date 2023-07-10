package RecodeTime

import (
	"testing"
	"time"
)

func TestRecodeTime(t *testing.T) {
	RecodeTime("start")
	time.Sleep(1 * time.Second)
	RecodeTime("first")
	time.Sleep(2 * time.Second)
	RecodeTime("second")
	time.Sleep(3 * time.Second)
	RecodeTime("thrid")

	PrintTime()
}
