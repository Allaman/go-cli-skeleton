package main

import (
	"io"
	"os"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readFromStdintoByte() ([]byte, error) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return stdin, nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func checkForEmptyByteSlice(data []byte) bool {
	return len(data) == 0
}
