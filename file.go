package goutil

import (
	"io/ioutil"
	"os"
)

// ListDirFiles is used to get all files in  one path
// map[string]interface{}
func ListDirFiles(dirPath string, suffix string) (interface{}, error) {
	files := make(map[string]interface{})

	listFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	pathSep := string(os.PathSeparator)
	tempFiles := make([]interface{}, 0, 10)
	for _, fi := range listFiles {
		tempPath := dirPath + pathSep + fi.Name()

		if fi.IsDir() {
			tempListFiles, _ := ListDirFiles(tempPath, suffix)
			tempFiles = append(tempFiles, tempListFiles)
			continue
		}

		tempFiles = append(tempFiles, fi.Name())
	}
	files[dirPath] = tempFiles
	return files, nil
}
