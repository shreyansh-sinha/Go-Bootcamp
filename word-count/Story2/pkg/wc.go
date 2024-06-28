package wordcount

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	common "example.com/common/utils"
)

func WordsCount(param string) {

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

	defer file.Close()

	// Initialize a word count variable
	wordCount := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into words
		words := strings.Fields(scanner.Text())
		// Increment the word count by the number of words in the line
		wordCount += len(words)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Print the total word count
	fmt.Printf("%8d %s\n", wordCount, param)
}
