package zos

import "os"

// Exist checks if a file or directory exists
func Exist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}