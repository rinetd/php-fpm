package main

import (
	"fmt"
	//	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

/*
go build  -ldflags "-s -w -H windowsgui" php-cgi.go
*/

func main() {
	//var php_process os.Process

	path := GetCurrPath()
	php_exe := path + "php-cgi.exe"
	php_ini := path + "php.ini"

	avg := []string{php_exe, "-b", "127.0.0.1:9090", "-c", php_ini}
	if _, err := os.Stat(php_exe); err != nil {
		return
	}

	for {
		/*
			cmd := exec.Command(php_exe, avg...)
			cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			//隐藏
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

			err := cmd.Start()
		*/

		procAttr := new(os.ProcAttr)
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		procAttr.Sys = &syscall.SysProcAttr{HideWindow: true}

		php_process, err := os.StartProcess(php_exe, avg, procAttr)

		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 10)
			continue
		}
		php_process.Wait()
		//		err = cmd.Wait()
		//fmt.Println("out")
		time.Sleep(time.Second * 1)
	}
}

/*获取当前文件执行的路径,,如c:/af/f/ */
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	splitstring := strings.Split(path, "\\")
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return ret
}
