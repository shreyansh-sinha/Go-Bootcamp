package wordcount

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	common "example.com/common/utils"
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
