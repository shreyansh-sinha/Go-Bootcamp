package wc

import (
	"bufio"
	"fmt"
	"os"
)

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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func LinesCount(param string) {

	if !fileExists(param) {
		fmt.Printf("'%s': open: No such file or Directory", param)
		return
	}

	if !isFile(param) {
		fmt.Printf("'%s': read: Is a directory", param)
		return
	}

	if !checkFileReadPermissions(param) {
		fmt.Printf("'%s': open: Permission denied", param)
		return
	}

	file, err := os.Open(param)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)

	defer file.Close()

	lineCount := 0

	for {
		_, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
	}
	fmt.Printf("%8d %s\n", lineCount, param)
}
