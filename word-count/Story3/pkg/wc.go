package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"os"

	common "example.com/common/utils"
)

func CharactersCount(param string) {

	res := common.Validate(param)

	if !res {
		return
	}

	file, err := os.Open(param)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Initialize a character count variable
	charCount := 0

	buffer := bufio.NewReader(file)
	data := make([]byte, 100)

	for {
		b, err := buffer.Read(data)
		if err == io.EOF {
			break
		}

		charCount += b
	}

	fmt.Printf("%8d %s\n", charCount, param)
}
