package wordcount

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	wc "example.com/Story3/pkg"
	"github.com/stretchr/testify/assert"
)

func TestWordCount_FileDoesNotExist(t *testing.T) {

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	old := os.Stdout
	os.Stdout = w

	defer func() { os.Stdout = old }()

	fileName := "test.txt"
	wc.CharactersCount(fileName)

	var buf bytes.Buffer
	w.Close()
	buf.ReadFrom(r)

	// Check the captured output against the expected error message
	expectedOutput := fmt.Sprintf("'%s': open: No such file or Directory", fileName)
	assert.Equal(t, expectedOutput, buf.String())

}

func TestWordCount_ParameterIsDirectory(t *testing.T) {

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	old := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = old }()
	defer func() { os.RemoveAll("cmd") }()

	directory := "cmd"

	// make directory to test directory read.
	err := os.Mkdir(directory, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	wc.CharactersCount(directory)

	// Capture the output from the pipe
	var buf bytes.Buffer
	w.Close()
	buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf("'%s': read: Is a directory", directory)
	assert.Equal(t, expectedOutput, buf.String())

}

func TestWordCount_FileProtected(t *testing.T) {

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	old := os.Stdout
	os.Stdout = w

	mode := os.FileMode(066) // read not allowed

	fileName := "file.txt"

	// Create or truncate the file with the specified permissions
	_, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fileName, err)
		return
	}

	defer func() { os.Remove(fileName) }()
	defer func() { os.Stdout = old }()

	wc.CharactersCount(fileName)

	// Capture the output from the pipe
	var buf bytes.Buffer
	w.Close()
	buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf("'%s': open: Permission denied", fileName)
	assert.Equal(t, expectedOutput, buf.String())

}

func TestWordsCount_FileWithContent_ReturnsLineCount(t *testing.T) {

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	// Redirect stdout to the pipe
	old := os.Stdout
	os.Stdout = w

	defer func() { os.Stdout = old }()

	/*
		create a temp file and write to it.
		once test execution is done, delete it.
	*/
	fileName := "test.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() { os.Remove(fileName) }()

	// write to test file.
	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte("This is Line 1.\nThis is Line 2"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	wc.CharactersCount(fileName)

	// Capture the output from the pipe
	var buff bytes.Buffer
	w.Close()
	buff.ReadFrom(r)

	expectedOutput := fmt.Sprintf("%8d %s\n", 30, fileName)
	assert.Equal(t, expectedOutput, buff.String())

}
