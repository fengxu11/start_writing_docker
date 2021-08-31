# Network Namespace

> Network Namespace 是用来隔离网络设备、IP地址端口等网络栈。

> Network Namespace 可以让每个容器拥有自己独立的虚拟网络设备、而且 容器内的应用可以绑定到自己的端口上、每个namespace中的端口都不会冲突、在宿主机搭建 网桥后 就能很方便的实现 容器间的通信

# Example

```sh

# 1. 查看当前宿主机的网络情况
[root@fengxu 5.user]# ifconfig
可以看到 lo、eth0、eth1 等网络设备

# 2. 运行 network namespace demo
[root@fengxu 5.user]# go run main.go
$ ifconfig
$ 

可以看到什么网络都没有

```

通过以上例子就可以看出、当前 namespace中的网络 和 宿主机的网络隔离开了