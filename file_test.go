package goutil

import (
	"io/ioutil"
	"testing"

	"os"

	"path/filepath"

	"sort"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListDirFiles(t *testing.T) {

	Convey("Test List Dir File", t, func() {
		pathName, _ := ioutil.TempDir("/tmp", "test")
		temp := []string{}

		expectResult := make(map[string]interface{})

		for i := 0; i < 10; i++ {
			filePath, _ := ioutil.TempFile(pathName, "test")
			temp = append(temp, filepath.Base(filePath.Name()))
		}
		sort.Strings(temp)
		tempInterface := []interface{}{}
		for _, i := range temp {
			tempInterface = append(tempInterface, i)
		}
		expectResult[pathName] = tempInterface

		Convey("Test Get Dir File List", func() {
			actualResult, _ := ListDirFiles(pathName, "")

			So(actualResult, ShouldResemble, expectResult)
		})
		Reset(func() {
			for _, name := range temp {
				os.Remove(name)
			}

			os.Remove(pathName)
		})
	})

}
