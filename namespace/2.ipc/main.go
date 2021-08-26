package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("sh")
	// 创建 IPC namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
