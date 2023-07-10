package RecodeTime

import (
	"fmt"
	"time"
)

type timeStamp struct {
	id        int
	timeStep  time.Time
	TimeSince time.Duration
	label     string
}

var TimeList []timeStamp

func RecodeTime(label string) {
	if len(TimeList) == 0 {
		TimeList = append(TimeList, timeStamp{
			id:        0,
			timeStep:  time.Now(),
			TimeSince: time.Duration(0),
			label:     label,
		})
		return
	}
	lastTime := TimeList[(len(TimeList) - 1)]
	TimeList = append(TimeList, timeStamp{
		id:        lastTime.id + 1,
		timeStep:  time.Now(),
		TimeSince: time.Since(lastTime.timeStep),
		label:     label,
	})
	return
}

func PrintTime() {
	fmt.Printf("%s \t %-30s \t\t %-10s \t\t %-10s \n", "id", "label", "cost time", "log time")
	for _, timeStamp := range TimeList {
		fmt.Printf("%d \t %-30s \t\t %-10f \t\t %-10s \n", timeStamp.id, timeStamp.label, timeStamp.TimeSince.Seconds(), timeStamp.timeStep)
	}
	TimeList = nil
}
