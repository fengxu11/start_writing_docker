# IPC
> IPC 全称是 Inter-Process Communication、指的是 多进程间通信的方法


## 先了解以下两个概念

> System V IPC 指的是在 System V.2发型版本中 引入的三种进程间通信工具 

> 信号量、共享内存、消息队列、我们把这三种工具统称为 System VIPC对象、每个对象都具有一个唯一的 IPC标识符

> POSIX message queues、消息队列、IPC通信的一种方式

## IPC namespace 就是用来隔离 System V IPC 和 POSIX message queues


## Example

> 1. 打开一个shell 使用ipcs工具 message mq
```sh
# 查看当前的消息队列
ipcs -q

------ Message Queues --------
key        msqid      owner      perms      used-bytes   messages    

# 创建一个 消息队列
ipcmk -Q

Message queue id: 0

# 查看当前的消息队列
ipcs -q

------ Message Queues --------
key        msqid      owner      perms      used-bytes   messages    
0x4bfaf52d 0          root       644        0            0    

```

> 2. 重新打开一个shell、进入到刚才的代码目录、运行代码 go run main.go
```sh

go run main.go
sh-4.4#

# 查看当前的消息队列

sh-4.4# ipcs -q

------ Message Queues --------
key        msqid      owner      perms      used-bytes   messages    


```

> 3. 可以看到、隔离后的环境下是看不到 宿主机上面创建的MQ的