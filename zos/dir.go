package zos

import "os"

// IsDir returns true if the given path is a directory.
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.Mode().IsDir()
}

// WorkDIR returns the current working directory.
func WorkDIR() string {
	w, _ := os.Getwd()
	return w
}