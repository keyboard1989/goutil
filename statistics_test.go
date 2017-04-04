package goutil

import (
	"strconv"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func MockGetTime() time.Time {
	return time.Now()
}

func TestStatistics(t *testing.T) {
	statis := NewStatistics(10000, 10, "2006-01-02 15:04:05", getTime)
	var i int
	i = 0
	for i < 100000 {
		item := Item{
			Type:    "api",
			SubType: "get",
			Key:     "user" + strconv.Itoa(i%19),
		}
		i += 1
		statis.AddItem(item)
	}

	time.Sleep(time.Second)
	info := statis.GetInfoByTypeAndSubType("api", "get")
	spew.Dump(info)
}
