package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

type DirElement struct {
	BaseName string
	Path     string
	IsDir    bool
	Ending   string
}

func readDir(path string, recursive bool) ([]DirElement, error) {
	var result []DirElement
	if recursive {
		err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				GL.logger.Debug().Msgf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			result = append(result, DirElement{BaseName: info.Name(), Path: filepath.Dir(path), IsDir: info.IsDir(), Ending: filepath.Ext(path)})
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		files, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			fsInfo, err := f.Info()
			if err != nil {
				GL.logger.Warn().Msgf("Issue while reading '%s'", f.Name())
			}
			result = append(result, DirElement{BaseName: fsInfo.Name(), Path: path, IsDir: fsInfo.IsDir(), Ending: filepath.Ext(f.Name())})
		}
	}
	return result, nil
}
