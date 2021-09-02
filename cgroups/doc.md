# Linux  Cgroups

> Cgroups的由来是 为了限制 一组进程及将来子进程 所能使用到的资源

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

