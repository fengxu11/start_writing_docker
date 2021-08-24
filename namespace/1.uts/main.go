package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 指定 fork出来的新进程内的初始命令、使用sh执行
	cmd := exec.Command("sh")
	// 创建一个 uts类型的 namespace、 go语言封装了对 clone函数的调用
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	// 指定标准输入、标准输出、标准错误输出
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
