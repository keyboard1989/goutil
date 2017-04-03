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
