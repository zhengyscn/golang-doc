package main

import (
    "fmt"
    "sync"
    "time"
)

/*
    读写锁： 读多写少
    互斥锁： 写多读少


    # 竞争检测
    $ go run -race main.go
    $ go build -race main.go
 */

var (
    m map[int]uint64
    lock sync.Mutex
)

type task struct {
    n int
}

func worker(t *task)  {

    var sum uint64 = 1
    for i:=1; i<t.n; i++ {
        sum *= uint64(i)
    }
    lock.Lock()
    m[t.n] = sum
    lock.Unlock()
}

func main()  {
    m = make(map[int]uint64)
    for i:=1; i<=16; i++ {
        t := &task{n:i}
        go worker(t)
    }

    time.Sleep(time.Second*10)
    lock.Lock()
    for k, v := range m {
        fmt.Printf("%d! = %d\n", k, v)
    }
    lock.Unlock()
}
