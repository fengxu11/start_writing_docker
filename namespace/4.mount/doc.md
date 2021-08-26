# Mount Namespace
> Mount Namespace 是用来隔离各个进程看到的挂载点视图、在不同的 namespace中看到的文件系统层次是不一样的

> 在 Mount Namespace中 调用 mount() 和 unmount() 仅会影响当前 namespace 内的文件系统、而全局的文件系统是没有影响的

> Mount Namespace 是Linux实现的第一个 namespace、因此系统调用参数是 NEWNS(New Namespace的缩写)

# Example

```sh

# 1. 运行代码
go run main.go

# 2. 查看一下 /proc的文件内容、这个时候 /proc还是宿主机的
sh-4.4# ls /proc
1     11019  13611  15441  16     17904  21     3     34   372  386   589   621  800   buddyinfo  dma          irq          loadavg  pagetypeinfo  swaps          vmallocinfo
10    11026  13632  15471  16527  18000  22     30    343  373  387   6     624  9     bus        driver       kallsyms     locks    partitions    sys            vmstat
100   11039  13709  15511  16861  18044  22361  31    345  379  4     6066  625  922   cgroups    execdomains  kcore        mdstat   sched_debug   sysrq-trigger  zoneinfo
101   11059  13798  15518  17     18473  23     3131  347  380  4206  611   628  926   cmdline    fb           keys         meminfo  schedstat     sysvipc
102   12     14     15557  17490  18507  25625  3196  349  381  48    614   686  929   consoles   filesystems  key-users    misc     scsi          thread-self
104   12013  15     15589  17619  18512  26     32    35   382  492   615   704  932   cpuinfo    fs           kmsg         modules  self          timer_list
1058  13     15246  15614  17620  18531  27     33    354  383  524   6158  732  933   crypto     interrupts   kpagecgroup  mounts   slabinfo      tty
1063  13591  15438  15721  17622  19     28     335   36   384  575   617   794  99    devices    iomem        kpagecount   mtrr     softirqs      uptime
11    13597  15440  15815  17816  2      29     339   37   385  587   6208  8    acpi  diskstats  ioports      kpageflags   net      stat          version

可以看到还是很乱的

# 3. 将文件系统挂在到当前 namespace中
# -t 指定文件系统的类型
# mount -t type(类型) device(设备) dir(目录)
mount -t proc proc /proc
mount: /proc: proc already mounted on /proc. (可以看到 已经已挂载)

# 4. 查看/proc 这个时候 /proc是命名空间内的
sh-4.4# ls /proc/
1          bus       cpuinfo    dma          filesystems  ioports   keys         kpagecount  mdstat   mounts        partitions   self      swaps          thread-self  version
4          cgroups   crypto     driver       fs           irq       key-users    kpageflags  meminfo  mtrr          sched_debug  slabinfo  sys            timer_list   vmallocinfo
acpi       cmdline   devices    execdomains  interrupts   kallsyms  kmsg         loadavg     misc     net           schedstat    softirqs  sysrq-trigger  tty          vmstat
buddyinfo  consoles  diskstats  fb           iomem        kcore     kpagecgroup  locks       modules  pagetypeinfo  scsi         stat      sysvipc        uptime       zoneinfo

# 5. 查看namespace中的 进程
sh-4.4# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 00:12 pts/5    00:00:00 sh
root         5     1  0 00:14 pts/5    00:00:00 ps -ef


```

通过以上例子就可以看出、当前 namespace中、sh进程 PID为1、
这说明 当前的 Mount namespace 中的 mount 和 外部空间是隔离的、
mount 操作并没有影响到外部、docker volume 也是利用了这个特性