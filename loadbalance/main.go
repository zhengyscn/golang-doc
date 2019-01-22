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

	fmt.Printf("instances:%v\n", instances)

	bal := balance.RandomBalance{}
	for {
		ins, err := bal.DoBalance(instances)
		fmt.Println(ins, err)
		time.Sleep(time.Second)
	}

}
