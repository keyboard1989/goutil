package goutil

import (
	"errors"
	"time"
)

// GetCurrentTimeStr is used to get current time str
// with format "2006-01-02 15:04:05"
func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetDayRange is used to get a range of start and end time
// start string eg:2017-01-02
// end   string eg:2017-01-22
// format string eg: 2006-01-02
// result: ["2017-01-02","2017-01-03",...,"2017-01-21"]
func GetDayRange(start string, end string, format string) ([]string, error) {
	ret := []string{}
	var err error

	startTime, err := time.Parse(format, start)
	if err != nil {
		return ret, err
	}

	endTime, err := time.Parse(format, end)
	if err != nil {
		return ret, err
	}

	if startTime.Unix() > endTime.Unix() {
		return ret, errors.New("start bigger than end")
	}

	for !startTime.Equal(endTime) {
		timeStr := startTime.Format(format)
		ret = append(ret, timeStr)
		startTime = startTime.Add(24 * time.Hour)
	}

	return ret, nil
}
