package goutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

// IsFileExist is used to check file exist or not
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDirExists is used to check dir is exist or not
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return fi.IsDir()
}

// IsAbsolutePath is used to check path is absolute path
func IsAbsolutePath(path string) bool {
	return strings.HasPrefix(path, string(rune(filepath.Separator)))
}
