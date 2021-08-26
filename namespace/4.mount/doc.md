# Mount Namespace
> Mount Namespace 是用来隔离各个进程看到的挂载点视图、在不同的 namespace中看到的文件系统层次是不一样的

> 在 Mount Namespace中 调用 mount() 和 unmount() 仅会影响当前 namespace 内的文件系统、而全局的文件系统是没有影响的

> Mount Namespace 是Linux实现的第一个 namespace、因此系统调用参数是 NEWNS(New Namespace的缩写)

# Example

