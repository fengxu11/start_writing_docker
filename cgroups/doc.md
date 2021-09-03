# Linux  Cgroups

> Cgroups (Control Groups) 的由来是 为了限制 一组进程及将来子进程 所能使用到的资源

> 这些资源包括 CPU、内存、存储、网络等、通过  Cgroups 可以很方便的限制某个进程的资源占用

## Cgroups 中的三个组件

### cgroup 

> Cgroup 是对进程分组管理的一种 机制，一个 cgroup包含一组进程、并可以在这个 cgroup上增加 Linux subsystem 的各种参数配置，将一组进程和一组 subsystem的系统参数关联起来


### subsystem
> 是一组 资源控制的模块，一般包含如下几项

1. blkio 设置对块设备（比如硬盘）输入输出的访问控制
2. cpu 设置cgroup 中进程的 CPU 被调度的策略
3. cpuacct 统计 cgroup中进程的CPU占用
4. cpuset 在多核机器上设置 cgroup中进程可以使用的 cpu 和 内存 (这里内存设置仅限于 NUMA架构)
5. devices 控制cgroup中进程对设备的访问
6. freezer     用于挂起(suspend) 和 恢复(resume) cgroup中的进程
7. memory      用于控制 cgroup中进程的内存占用
8. net_cls      用于将 cgroup 中进程产生的网络包分类、以便于 linux的 tc(traffic controller) 可以根据分类区分出 来自某个cgroup的包、并做限流或者监控
9. net_prio     设置cgroup中 进程产生的网络流量的优先级
10. ns          这个subsystem 比较特殊、它的作用是使 cgroup中的进程 在新的namespace中 fork新进程(NEWNS)时、创建一个新的cgroup、这个cgroup包含新的namespace中的进程

> 每个 subsystem会关联到 定义了相应限制的cgroup上、并对这个cgroup中的进程做出 相应的限制和控制、

### 这些 subsystem 是逐步合并到内核中的、如何查看当前的内核支持哪些 subsystem呢 ？

> 安装相应的 cgroup 工具、然后通过lssubsys -a 看到 kernel支持的 subsystem 
```
apt-get install cgroup-bin
```


### hierarchy
> hierarchy 的功能 是把一组 cgroup 串成一种树状的结构、一个这样的树便是一个 hierarchy、
> 通过 hierarchy cgroups可以做到继承

> 比如、系统对一组定时任务的进程通过 cgroup1 限制了CPU的使用率、然后其中有一个定时任务(一个进程)还需要限制
> 磁盘IO、为了避免限制磁盘io 影响到了其他进程、这时创建一个 cgroup2 让其继承于cgroup1、这样 cgroup2
> 使其继承于 cgroup1的 CPU使用率限制、并且增加了磁盘IO的限制、而不影响cgroup1中的其他进程


### 三个组件之间的关系

1. 系统创建了 hierarchy之后、系统中所有进程都会加入到 这个 hierarchy中的 cgroup根节点上、这个根节点是 hierarchy默认创建的
2. 一个 subsystem 只能附加到一个 hierarchy
3. 一个 hierarchy可以让 多个 subsystem附加上
4. 一个进程 可以作为多个cgroup的成员、但是这些cgroup必须在 不同的 hierarchy中
5. 一个进程 fork出子进程时、子进程和父进程在同一个cgroup中的、可以根据需要将子进程 移动到其他cgroup中的



### Kernel接口、怎么调用 kernel才能配置 Cgroups呢？

> 前面说过 cgroup中的 hierarchy是一种树状结构、kernel为了对cgroups的配置更加直观、
> 是通过 一个虚拟的树状文件系统配置 Cgroups的、通过层级目录虚拟出cgroup树。

> 列出当前系统支持哪些 subsystem、如果没有 lssubsystem这个工具、可以使用这个命令安装
>  ```yum install libcgroup-tools
```
[root@fengxu fengxu]# lssubsys -a
cpuset
cpu,cpuacct
blkio
memory
devices
freezer
net_cls,net_prio
perf_event
hugetlb
pids
rdma
```

1. 首先、创建并挂在一个 hierarchy(cgroup树) 

```

# 创建一个 hierarchy 挂载点
[root@fengxu]# mkdir cgroup-test

# 挂载一个 hierarchy
[root@fengxu]# sudo mount -t cgroup  -o none,name=cgroup-test cgroup-test  ./cgroup-test

# 挂在后 就可以看到系统在这个目录下生成了、一些默认文件
[root@fengxu]# ls ./cgroup-test
cgroup.clone_children  cgroup.procs  cgroup.sane_behavior  notify_on_release  release_agent  tasks


```
> cgroup.clone_clhldren       cpuset的 subsystem 会读取这个配置文件、如果这个值是 1(默认是0)、
> 子 cgroup才会继承父cgroup的cpuset配置

> cgroup.procs 是树的当前节点cgroup中的进程组ID、现在的位置是在根节点、这个文件中会有现在系统中所有进程组的ID

> notify_on_release 和 release_agent 会一起使用。  notify_on_release 标识当这个cgroup 最后一个进程退出的时候是否执行了 release_agent; release_agent 则是一个路径，通常用作进程退出之后自动清理掉不再使用 的cgroup

> tasks 标识该cgroup 中的进程ID，如果把一个进程ID 写入到tasks文件中、便会将相应的进程加入到这个cgroup中



2. 通过刚创建好的 hierarchy上 cgroup根节点中 扩展出的两个子 cgroup

```
# 在 刚创建好的 hierarchy上的cgroup根节点中创建两个 cgroup (也就是 在cgroup-test文件夹中创建)

[root@fengxu cgroup-test]# sudo mkdir cgroup-1
[root@fengxu cgroup-test]# sudo mkdir cgroup-2
[root@fengxu cgroup-test]# tree
.
├── cgroup-1
│   ├── cgroup.clone_children
│   ├── cgroup.procs
│   ├── notify_on_release
│   └── tasks
├── cgroup-2
│   ├── cgroup.clone_children
│   ├── cgroup.procs
│   ├── notify_on_release
│   └── tasks
├── cgroup.clone_children
├── cgroup.procs
├── cgroup.sane_behavior
├── notify_on_release
├── release_agent
└── tasks

2 directories, 14 files


# 可以看到、在一个 cgroup下创建文件夹 时、kernel会把文件夹标记为 这个cgroup的子cgroup、并继承他们的父属性
```


3. 在cgroup中 添加 和 移动进程

```

[root@fengxu cgroup-test]# echo $$
18059

# 进入到 cgroup-1中
[root@fengxu cgroup-test]# cd cgroup-1/
[root@fengxu cgroup-1]# ls
cgroup.clone_children  cgroup.procs  notify_on_release  tasks

# 将 当前所在的终端进程移动到 cgroup-1中
[root@fengxu cgroup-1]# sudo sh -c "echo $$ >> tasks"


# 查看终端进程 所在的 cgroup
[root@fengxu cgroup-1]# cat /proc/18059/cgroup 
13:name=cgroup-test:/cgroup-1
12:rdma:/
11:devices:/user.slice
10:cpu,cpuacct:/user.slice
9:blkio:/user.slice
8:net_cls,net_prio:/
7:perf_event:/
6:hugetlb:/
5:freezer:/
4:memory:/user.slice/user-0.slice/session-10.scope
3:cpuset:/
2:pids:/user.slice/user-0.slice/session-10.scope
1:name=systemd:/user.slice/user-0.slice/session-10.scope

```


4. 通过 subsystem限制 cgroup中进程的资源

```

# 在上面创建hierarchy的时候、并没有关联任何的 subsystem，所以没办法通过那个 hierarchy中的 cgroup节点限制进程的资源占用

# 系统默认已经为每个 subsystem创建了一个hierarchy、比如  memory的hierarchy、查看一下
# 可以看到 /sys/fs/cgrooup/memory 目录已经挂载到 memory subsystem的 hierarchy上面
[root@fengxu cgroup-1]# mount | grep memory
cgroup on /sys/fs/cgroup/memory type cgroup (rw,nosuid,nodev,noexec,relatime,memory)

# 接下来的例子 就在这个 hierarchy中 创建cgroup、限制进程所使用的cpu

# 1. 进入系统CPU的 hierarchy中、创建一个 cgroup
[root@fengxu]# cd /sys/fs/cgroup/cpu

[root@fengxu]# mkdir cgroups_test

[root@fengxu cgroups_test]# ls
cgroup.clone_children  cpuacct.usage_all          cpuacct.usage_sys   cpu.rt_period_us   notify_on_release
cgroup.procs           cpuacct.usage_percpu       cpuacct.usage_user  cpu.rt_runtime_us  tasks
cpuacct.stat           cpuacct.usage_percpu_sys   cpu.cfs_period_us   cpu.shares
cpuacct.usage          cpuacct.usage_percpu_user  cpu.cfs_quota_us    cpu.stat

# 2. 设置这个cgroup中的进程 只能使用20%的系统CPU资源
[root@fengxu cgroups_test]# echo 20000 > cpu.cfs_quota_us

# 3. 启动一个新的bash、运行一个程序、打满cpu
[root@fengxu cgroups_test]# while : ; do : ; done &
[1] 23682

# 4. 把这个进程id、加入到这个cgroup中 
[root@fengxu cgroups_test]# echo 23682 > tasks


# 5. 使用 top 查看进程占用的cpu
  PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                      
23682 root      20   0   26644   3128   1260 R  19.9   0.0   0:22.06 bash  

可以看到 23682这个进程 CPU使用不会超过 20%

为什么 20000代表的是 CPU使用率是20%呢？
cpu.cfs_quota_us  表示cgroup限制占用的时间、单位是 微妙、默认为-1、
设置为20000表示 20000/10000=20%的CPU

```



### docker 是如何使用 Cgroups 的？

> 先安装好docker 

1. 启动一个nginx容器、并进入
# docker run -it --cpus=".5" nginx /bin/sh
# ls
bin   dev                  docker-entrypoint.sh  home  lib64  mnt  proc  run   srv  tmp  var
boot  docker-entrypoint.d  etc                   lib   media  opt  root  sbin  sys  usr

2. 进入cpu的cgroup中
# cd /sys/fs/cgroup/cpu  
# ls
cgroup.clone_children  cpu.rt_period_us   cpuacct.stat          cpuacct.usage_percpu_sys   notify_on_release
cgroup.procs           cpu.rt_runtime_us  cpuacct.usage         cpuacct.usage_percpu_user  tasks
cpu.cfs_period_us      cpu.shares         cpuacct.usage_all     cpuacct.usage_sys
cpu.cfs_quota_us       cpu.stat           cpuacct.usage_percpu  cpuacct.usage_user

3. 查看CPU限制
# cat cpu.cfs_quota_us
50000


### Go语言实现 限制进程的CPU资源

1. 运行 main.go
```
[root@fengxu cgroups]# go run main.go 
```


2. 重新打开一个bash、使用top查看 刚才 shell 进程、可以看到 CPU没有超过30%

```
  PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                
 8873 root      20   0   12700    252      0 R  30.0   0.0   0:02.96 sh         
```

3. 停止main.go之后、记得清除cgroup

```
[root@fengxu cpu]# cgdelete cpu:test_limit_cpu
```