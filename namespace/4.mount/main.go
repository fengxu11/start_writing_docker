package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Mount Namespace 是Linux实现的第一个 namespace、因此系统调用参数是 NEWNS(New Namespace的缩写)
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
