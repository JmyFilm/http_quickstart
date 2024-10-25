package utils

import (
	"os"
	"path/filepath"
)

func RunPath() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}
