package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zhengyscn/golang-doc/loadbalance/balance"
)

func main() {
	instances := make([]*balance.Instnace, 0)

	for i := 0; i < 3; i++ {
		host := fmt.Sprintf("192.168.%d,%d", rand.Intn(23), rand.Intn(23))
		inst := balance.NewInstance(host, 8080)
		instances = append(instances, inst)
	}

	var randomBalance = balance.RoundrobinBalance{}
	var doBalance balance.Balance

	// interface 赋值
	doBalance = &randomBalance // 是否通过指针赋值，取决与randomBalance结构体下的函数

	for {
		ins, err := doBalance.DoBalance(instances)
		if err != nil {
			fmt.Errorf("Do balance err:%v\n", err)
			break
		}
		fmt.Println(ins)
		time.Sleep(time.Second)
	}

}
