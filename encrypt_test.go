package goutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMd5(t *testing.T) {
	Convey("TestMd5", t, func() {
		originStr := "abc"
		expectStr := "900150983cd24fb0d6963f7d28e17f72"

		actualSt := Md5(originStr)
		So(actualSt, ShouldEqual, expectStr)
	})
}

func TestBase64(t *testing.T) {
	Convey("Test Base64 Encode", t, func() {
		Convey("Test", func() {
			originStr := "abc"
			expectStr := "YWJj"
			actualStr := Base64Encode(originStr)

			So(actualStr, ShouldEqual, expectStr)
		})

		Convey("Test Base64 decode", func() {
			originStr := "YWJj"
			expectStr := "abc"

			actual, err := Base64Decode(originStr)

			So(err, ShouldBeNil)
			So(string(actual), ShouldEqual, expectStr)
		})
	})
}
