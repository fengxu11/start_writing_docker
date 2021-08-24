
# namespace uts

> 创建一个 uts类型的 namespace

> uts 类型的namespace、主要用来隔离 nodename 和 domainname两个标识

> 在 uts类型的namespace中、允许每个namespace有自己的 hostname

# 测试 uts namespace 流程

1. 进入 uts文件夹、运行 go run main.go
2. 会进入到一个交互式的环境中、使用 pstree -pl 查看整个linux系统中 进程之间的关系
3. 查看 当前这个sh进程的 PID、echo $$
4. 查看父子进程 是否在同一个namespace中、父进程指的是 go run main.go 运行起来的程序、而子进程指的是 sh这个进程
```bash
6373指的是父进程
readlink /proc/6373/ns/uts
6377指的是子进程
readlink /proc/6377/ns/uts

# 结果
sh-4.4# readlink /proc/6373/ns/uts
uts:[4026531838]
sh-4.4# readlink /proc/6377/ns/uts
uts:[4026532349]
sh-4.4# 
```

5. 结果很明显 两个进程并不在一个 namespace中
6. 接下来测试一下修改namespace中的 hostname 会不会影响到外部主机的
```sh
# 在sh环境下 修改hostname、然后在打印hostname
sh-4.4# hostname -b test-host-name
sh-4.4# hostname
test-host-name

# 另外打开一个shell、查看外部主机的hostname
sh-4.4# hostname
iZ2lrgfg4v5709Z
```
7. 经 第六步实验可以发现、uts namespace中 hostname标识修改后、不影响外部主机

