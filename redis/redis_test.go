package redis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConnect(t *testing.T) {

	Convey("test redis connect", t, func() {
		address := "127.0.0.1:6379"
		testRedis, _ := NewRedis(address)

		Convey("test get", func() {
			expectName := "zhangyong"
			testRedis.Set("name", expectName)
			name, _ := testRedis.Get("name")
			So(name, ShouldEqual, expectName)
		})

		Convey("test set", func() {
			testKey := "test_key"
			expectName := "test set string"

			testRedis.Set(testKey, expectName)

			actual, _ := testRedis.Get(testKey)

			So(expectName, ShouldEqual, actual)
		})

		Reset(func() {
			testRedis.Close()
		})

	})

}
