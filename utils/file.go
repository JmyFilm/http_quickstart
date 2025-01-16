package utils

import (
	"os"
	"path/filepath"
)

func NewFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	fp, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Close()
}
