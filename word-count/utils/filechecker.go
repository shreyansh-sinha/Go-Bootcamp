package common

import (
	"fmt"
	"os"
)

// check if there exists a file with given parameter
func IsFile(param string) bool {

	fileInfo, err := os.Stat(param)
	if err != nil {
		return false
	}

	if fileInfo.Mode().IsRegular() {
		return true
	}

	return false

}

// check if file with given parameter has read permissions
func CheckFileReadPermissions(param string) bool {

	fileInfo, err := os.Stat(param)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if fileInfo.Mode().Perm()&0400 != 0 {
		return true
	}

	return false
}

// check if file or directory exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
