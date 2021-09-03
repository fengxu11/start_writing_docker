package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"
)

const procSelf = "/proc/self/exe"
const cgroup = "test_limit_cpu"
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/cpu"

func main() {
	if os.Args[0] == procSelf {
		// cmd进程在cgroup规则下运行
		cmd := exec.Command("sh", "-c", "while : ; do : ; done &")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}

		// 阻塞
		for {
			time.Sleep(1 * time.Second)
		}
	}

	cmd := exec.Command(procSelf)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 启动 一个进程
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	fmt.Println("命名空间外部的进程ID: ", cmd.Process.Pid)
	// 创建cgourp
	os.Mkdir(path.Join(cgroupMemoryHierarchyMount, cgroup), 0755)
	// 把当前进程加入到 cgroup中、并且限制 cpu使用率不能超过 30%
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, cgroup, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, cgroup, "cpu.cfs_quota_us"), []byte("30000"), 0644)

	// 接收命令 输出结果
	cmd.Process.Wait()
}
