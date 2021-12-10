package zos

import (
	"os"
	"path/filepath"
)

// OpenFile opens a file if it exists, otherwise creates it and directory.
func OpenFile(path string, perm os.FileMode) (*os.File, error) {
	dir, _ := filepath.Split(path)
	if !Exist(dir) {
		os.MkdirAll(dir, perm)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	return f, err
}