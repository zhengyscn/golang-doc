package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	多个goroutine 同时修改map，会引起资源竞争
*/

func main() {
	var m map[int]int = make(map[int]int, 5)
	m[2] = 1
	m[3] = 2

	fmt.Printf("before m :%v\n", m)

	for i := 0; i < 2; i++ {
		go func() {
			m[2] = rand.Intn(20)
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("after m :%v\n", m)
}
