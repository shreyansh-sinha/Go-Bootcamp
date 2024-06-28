package wc

import (
	"bufio"
	"fmt"
	"os"

	common "example.com/common/utils"
)

func LinesCount(param string) {

	if !common.FileExists(param) {
		fmt.Printf("'%s': open: No such file or Directory", param)
		return
	}

	if !common.IsFile(param) {
		fmt.Printf("'%s': read: Is a directory", param)
		return
	}

	if !common.CheckFileReadPermissions(param) {
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
