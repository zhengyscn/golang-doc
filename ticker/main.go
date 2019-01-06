package main

import (
	"fmt"
	"time"
)

func t1() {
	t := time.NewTicker(time.Second * 2)
	/*
	   t.C  -> goroutine 定时向channel写数据
	*/
	for v := range t.C {
		y, m, d := v.Date()
		date := fmt.Sprintf("%v-%v-%v", y, m, d)
		fmt.Printf("v = %v\n", date)
	}
	t.Stop()
}

func t2() {
	for {
		select {
		case <-time.After(time.Second * 2): // 不建议使用，linux有内存泄漏.
			fmt.Println("after 2 second.")
		}
	}

}

func main() {
    t1()
	t2()
}
