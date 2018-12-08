# etcd

- 基于raft算法的强一致性
- 配置管理、服务发现


```
etcd-v3.3.10-linux-amd64.tar.gz

$ ETCD_VER=v3.3.10

$ ./etcd --version
etcd Version: 3.3.10
Git SHA: 27fc7e2
Go Version: go1.10.4
Go OS/Arch: linux/amd64


$ etcd --config-file=conf/conf.yml or etcd --data-dir=/tmp/data/

$ etcdctl put /message hello
$ etcdctl get /message
$ etcdctl del /message

注意： 使用go.etcd.io/etcd/clientv3这个包
```


