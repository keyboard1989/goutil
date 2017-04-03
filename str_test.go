package goutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubStr(t *testing.T) {
	Convey("TestSubStr", t, func() {

		Convey("Test Normal", func() {
			originStr := "abcdefghijklmnopqrst"
			expectStr := "abcde"
			actualStr := SubStr(originStr, 0, 5)
			So(actualStr, ShouldEqual, expectStr)

			expectStr = "bcde"
			actualStr = SubStr(originStr, 1, 5)
			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test start and end too long", func() {
			originStr := "abcdefghijklmnopqrst"
			expectStr := ""
			actualStr := SubStr(originStr, 100, 500)
			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test end too long", func() {
			originStr := "abcdefghijklmnopqrst"
			expectStr := "hijklmnopqrst"
			actualStr := SubStr(originStr, 7, 500)
			So(actualStr, ShouldEqual, expectStr)
		})
	})
}

func TestSubStrByOffset(t *testing.T) {
	Convey("TestSubStrByOffset", t, func() {

		Convey("Test Normal", func() {
			originStr := "abcdefghijklmnopqrst"
			expectStr := "abcde"
			actualStr := SubStrByOffset(originStr, 0, 5)
			So(actualStr, ShouldEqual, expectStr)

			expectStr = "bcdef"
			actualStr = SubStrByOffset(originStr, 1, 5)
			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test start too long", func() {
			originStr := "abcdefghijklmnopqrst"
			expectStr := ""

			actualStr := SubStrByOffset(originStr, 100, 10)
			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test offset too long", func() {
			originStr := "0123456789"
			expectStr := "89"

			actualStr := SubStrByOffset(originStr, 8, 10)
			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test unicode", func() {
			originStr := "01热烈欢迎23456789"
			expectStr := "1热烈"

			actualStr := SubStrByOffset(originStr, 1, 3)
			So(actualStr, ShouldEqual, expectStr)
		})
	})
}
