package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	互拆锁 和 读写锁 对比
*/

func main() {
	var lock sync.Mutex
	//var rwlock sync.RWMutex
	var count int32
	/*
	   100个写操作
	   1000个读操作
	*/

	var m map[int]int = make(map[int]int, 5)
	m[2] = 1
	m[3] = 2

	for i := 0; i < 100; i++ {
		go func() {
			lock.Lock() // 写操作加锁
			m[2] = rand.Intn(20)
			lock.Unlock()
		}()
	}
	for i := 0; i < 5000; i++ {
		go func() {
			for {
				lock.Lock()
				//fmt.Println(m)
				lock.Unlock()
				atomic.AddInt32(&count, 1) // 原子操作
			}
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}
