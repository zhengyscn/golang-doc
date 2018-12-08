package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zhengyscn/golang-doc/etcd/etcd"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		endpoints   = []string{"localhost:2379"}
		dialtimeout = time.Second * 5
	)
	client, err := etcd.NewClient(endpoints, dialtimeout)
	if err != nil {
		fmt.Printf("New etcd client failed, err:%v\n", err)
		return
	}
	defer client.Close()

	_, err = client.Put(context.Background(), "abc", "123456")
	if err != nil {
		fmt.Printf("Put failed, err:%v\n", err)
		return
	}

	resp, err := client.Get(context.Background(), "abc")
	if err != nil {
		fmt.Printf("Get failed, err:%v\n", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("k:%s, v:%s\n", kv.Key, kv.Value)
	}

	respChan := client.Watch(context.Background(), "/zhengyscn/192.168.1.100/config", clientv3.WithPrefix())
	for {
		select {
		case msg := <-respChan:
			for _, ev := range msg.Events {
				fmt.Printf("watch key:%s, val:%s\n", ev.Kv.Key, ev.Kv.Value)
			}
		}
	}

}
