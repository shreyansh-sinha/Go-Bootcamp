package common

import (
	"fmt"
	"os"
)

// check if there exists a file with given parameter
func isFile(param string) bool {

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
func checkFileReadPermissions(param string) bool {

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
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// validate file
func Validate(param string) bool {

	if !fileExists(param) {
		fmt.Printf("'%s': open: No such file or Directory", param)
		return false
	}

	if !isFile(param) {
		fmt.Printf("'%s': read: Is a directory", param)
		return false
	}

	if !checkFileReadPermissions(param) {
		fmt.Printf("'%s': open: Permission denied", param)
		return false
	}

	return true
}
