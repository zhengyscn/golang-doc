package main

import (
	"fmt"
	"time"

	"github.com/etcd-io/etcd/clientv3"
	"golang.org/x/net/context"
)

var cli *clientv3.Client

func NewClient(endpoints []string, dialtimeout time.Duration) (cli *clientv3.Client, err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialtimeout,
	})
	return
}

func Put(key string, val string) (err error) {
	resp, err := cli.Put(context.Background(), key, val)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func Del() {

}

func ClientClose() {
	cli.Close()
}

func main() {
	var (
		err         error
		endpoints   = []string{"localhost:2379"}
		dialtimeout = time.Second * 5
	)
	cli, err = NewClient(endpoints, dialtimeout)
	fmt.Println(cli)
	fmt.Println(err)

	ClientClose()
}
