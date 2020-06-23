package os_u

import (
	"os"
	"path/filepath"
)

var Exports = map[string]interface{}{
	"_name":     "os",
	"args":      os.Args[1:],
	"stdin":     os.Stdin,
	"stderr":    os.Stderr,
	"stdout":    os.Stdout,
	"getenv":    os.Getenv,
	"open":      os.Open,
	"create":    os.Create,
	"exit":      os.Exit,

	"Args":   os.Args[1:],
	"Stdin":  os.Stdin,
	"Stderr": os.Stderr,
	"Stdout": os.Stdout,
	"Getenv": os.Getenv,

	"Open":   os.Open,
	"Create": os.Create,
	"Exit":   os.Exit,
	"Pid": os.Getpid(),
}


// 程序工作目录
func WorkDIR() string {
	w, _ := os.Getwd()
	return w
}

// 打开文件
// 不存在则创建对应的文件和路径
func OpenFile(path string, perm os.FileMode) (*os.File, error) {
	dir, _ := filepath.Split(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, perm)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	return f, err
}

func Exit() {
	os.Exit(0)
}

func SafeExit() {
	panic("exit")
}

// 文件是否存在
func FileExists(file string) (exist bool) {
	_, err := os.Stat(file)
	exist = !os.IsNotExist(err)
	return
}

// 是否是目录
func IsDir(file string) (exist bool, err error) {
	f, err := os.Stat(file)
	if err != nil {
		return
	}
	exist = f.Mode().IsDir()
	return
}