# go module

## Golang1.12

```go
$ go version
go version go1.12.6 darwin/amd64
```

## 初始化一个module
```go
$ go mod init <xgin> 
```

- go.mod 会记录项目所需依赖以及具体的版本号记录
- go.sum 会记录依赖树的详细信息

## 其它命令行参数
```go
$ go help mod
    download    download modules to local cache            # 下载模块包到本地缓存
    edit        edit go.mod from tools or scripts          # 从工具或脚本中修改go.md内容(暂时没觉得很好用)
    graph       print module requirement graph             # 打印模块依赖树
    init        initialize new module in current directory # 在当前目录初始化一个模块
    tidy        add missing and remove unused modules      # 添加或略或删除不使用的包
    vendor      make vendored copy of dependencies         # 将依赖拷贝到vendor中(会将依赖打包到当前目录下的vendor目录)
    verify      verify dependencies have expected content  # 检查依赖内容
    why         explain why packages or modules are needed # 解释为什么哪些包或者模块被需要

```

## vendor拷贝
```go
$ go mod vendor  // 会将项目的依赖包拷贝到项目的vendor目录中
```


## goproxy

- https://goproxy.cn : 七牛云赞助支持的开源代理
- https://mirrors.aliyun.com/goproxy/ : 阿里云官方维护的go代理
- https://goproxy.io/ : 也是一个开源的go代理

```go
# 设置代理(上述三个代理地址都可以)
$ export GOPROXY=https://goproxy.cn


```