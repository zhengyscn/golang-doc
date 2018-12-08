package etcd

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	Prefix = clientv3.WithPrefix()
)

func NewClient(endpoints []string, dialtimeout time.Duration) (cli *clientv3.Client, err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialtimeout,
	})
	return
}
