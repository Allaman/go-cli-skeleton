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
	return dat, nil
}

func writeFile(data []byte, path string) error {
	exist, err := exists(path)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("file '%s' exists", path)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return err
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
