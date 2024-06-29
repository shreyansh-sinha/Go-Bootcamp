package wordcount

import (
	"bufio"
	"fmt"
	"os"

	common "example.com/common/utils"
)

func LinesCount(param string) {

	res := common.Validate(param)

	if !res {
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
