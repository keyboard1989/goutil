package goutil

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCurrentTimeStr(t *testing.T) {
	expectStr := time.Now().Format("2006-01-02 15:04:05")
	actualStr := GetCurrentTimeStr()

	if expectStr != actualStr {
		t.Errorf("GetCurrentTimeStr() func error,actual:%s expect:%s", actualStr, expectStr)
	}
}

func TestGetDayRange(t *testing.T) {
	Convey("Test Get day Range", t, func() {

		Convey("Test right situation", func() {
			start := "2017-01-02"
			end := "2017-01-05"
			format := "2006-01-02"
			expectResult := []string{"2017-01-02", "2017-01-03", "2017-01-04"}

			actualResult, err := GetDayRange(start, end, format)

			So(err, ShouldBeNil)
			So(actualResult, ShouldResemble, expectResult)
		})

		Convey("Test format error", func() {
			start := "2017-01-02"
			end := "2017-01-05"
			format := "200601-02"
			expectResult := []string{}

			actualResult, err := GetDayRange(start, end, format)

			So(err, ShouldNotBeNil)
			So(actualResult, ShouldResemble, expectResult)
		})

		Convey("Test start format error", func() {
			start := "201701-02"
			end := "2017-01-05"
			format := "2006-01-02"
			expectResult := []string{}

			actualResult, err := GetDayRange(start, end, format)

			So(err, ShouldNotBeNil)
			So(actualResult, ShouldResemble, expectResult)
		})

		Convey("Test end format error", func() {
			start := "2017-01-02"
			end := "201701-05"
			format := "2006-01-02"
			expectResult := []string{}

			actualResult, err := GetDayRange(start, end, format)

			So(err, ShouldNotBeNil)
			So(actualResult, ShouldResemble, expectResult)
		})

		Convey("Test start bigger than end", func() {
			start := "2017-02-02"
			end := "2017-01-05"
			format := "2006-01-02"
			expectResult := []string{}

			actualResult, err := GetDayRange(start, end, format)

			So(err, ShouldNotBeNil)
			So(actualResult, ShouldResemble, expectResult)
		})
	})
}
