package goutil

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func MockGetTime() time.Time {
	t1, _ := time.Parse("2006-01-02 15:04:05", "2017-08-29 12:12:00")
	return t1
}

func TestStatistics(t *testing.T) {

	Convey("Test statistics", t, func() {
		statis := NewStatistics(10000, 10, "2006-01-02 15:04:05", MockGetTime)
		var i int
		i = 0
		for i < 1000 {
			item := Item{
				Type:    "api",
				SubType: "get",
				Key:     "user",
			}
			i += 1
			statis.AddItem(item)
		}

		time.Sleep(time.Second)
		expectStr := `[{"time":"2017-08-29 12:12:00","user":1000}]`

		info := statis.GetInfoByTypeAndSubType("api", "get")
		actualJson, _ := json.Marshal(info)

		So(string(actualJson), ShouldEqual, expectStr)
	})
}
