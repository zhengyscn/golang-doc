package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	var m map[int]int = make(map[int]int, 5)
	m[2] = 1
	m[3] = 2

	fmt.Printf("before m :%v\n", m)
	for i := 0; i < 2; i++ {
		go func() {
			lock.Lock() // 写操作加锁
			m[2] = rand.Intn(20)
			lock.Unlock()
		}()
	}
	time.Sleep(time.Second)
	lock.Lock() // 读作加锁
	fmt.Printf("after m :%v\n", m)
	lock.Unlock()
}
