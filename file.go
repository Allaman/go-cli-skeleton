package main

import (
	"fmt"
	"os"
)

func readFile(path string) ([]byte, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if checkForEmptyByteSlice(dat) {
		GL.logger.Info().Msg("reading empty file")
	}
	return dat, nil
}

// writeBytesToFile creates a new file and writes content
// provided by parameter or stdin
// When the file exists it aborts
// When the content is empty the created file is also empty
func writeBytesToFile(data []byte, path string) error {
	// TODO: make this better
	var err error
	var content []byte
	// FIX: this is not good design. A caller does not expect that the
	// FIX: provdied data can be overwritten
	if isInputFromPipe() {
		GL.logger.Info().Msg("reading from pipe")
		content, err = readFromStdintoByte()
		if err != nil {
			return err
		}
	} else {
		content = data
	}
	exist, err := exists(path)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("file '%s' exists", path)
	}
	if checkForEmptyByteSlice(content) {
		GL.logger.Info().Msg("writing empty file - use stdin or flag for providing content")
	}
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return err
}
