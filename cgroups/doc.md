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

1. 首先、创建并挂在一个 hierarchy(cgroup树) 

2. 通过刚创建好的 hierarchy上 cgroup根节点中 扩展出的两个子 cgroup

3. 在cgroup中 添加 和 移动进程

4. 通过 subsystem限制 cgroup中进程的资源


### docker 是如何使用 Cgroups 的？


### 用Go语言实现 通过Cgroups 限制容器的资源