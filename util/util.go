package util

import (
	"io/ioutil"
)

// ReadDiskFile reads a file from disk as a string
func ReadDiskFile(path string) (fileContents string, err error) {
	fileBytes, err := ioutil.ReadFile(path)
	return string(fileBytes), err
}
