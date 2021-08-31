# User Namespace
> User Namespace 主要是用来隔离 用户和用户组ID的、

> 也就是说一个进程的 UID 和 GID 在 user namespace 中是不同的

> linux kernel 3.8开始 非root进程也可以创建 user namespace、并且此用户在namespace中 可以被映射成root、而且在 namespace中有 root权限

# CentOS 默认没有开启 user namespace、所以需要开启 user namespace

```
1. 设置
# grubby --args="user_namespace.enable=1" --update-kernel="$(grubby --default-kernel)"

2. 配置最大的 user namespace
echo "user.max_user_namespaces=10000" >> /etc/sysctl.conf

3. 重启服务器
reboot
```


# Example

```sh

# 1. 查看当前用户和用户组
[root@fengxu 5.user]# id
uid=0(root) gid=0(root) groups=0(root)
可以看到是  root用户

# 2. 运行 user namespace demo、进入shell后查看 id
sh-4.4$ id
uid=1234 gid=1234 groups=1234

```

通过以上例子就可以看出、他们的UID 和 GID是不同的、说明 user namespace生效了