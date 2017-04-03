package mysql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Configuration struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Type  int    `db:"type"`
	Value string `db:"value"`
}

func TestOpen(t *testing.T) {
	Convey("Test mysql", t, func() {
		dataSource := "root:@tcp(127.0.0.1:3306)/test?parseTime=true&charset=utf8&loc=Local"

		testDB := NewMysql(dataSource)
		config := []Configuration{}
		testDB.Query(&config, "select * from configuration limit 1")
		So(len(config), ShouldEqual, 1)
	})
}
