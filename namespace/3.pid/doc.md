# PID
> PID namespace 是用来隔离进程 id的

> 这样就明白了 为什么在 docker container 里面和外面pid不一样了

## Example

> 1. 打开一个shell 
```sh
# 进入shell中

go run main.go

sh-4.4# 

# 输入 echo $$ 输出 pid是1
sh-4.4# echo $$
1
 
# 重新打开一个 shell窗口、查看进程树、可以看到进程id是 16632
pstree -pl

sshd(922)─┬─main(16632)─┬─sh(16636)

```

通过以上例子就可以看出 container 里面和外面pid不一样、这样就验证了 PID 隔离了