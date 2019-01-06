package main

import (
	"fmt"
	"time"
)

func worker(ch1, ch2 chan int) {
	var i int
	for {
		ch1 <- i
		time.Sleep(time.Second)
		ch2 <- i * i
		i++
	}
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go worker(ch1, ch2)

	// channel 不阻塞
	for {
		select {
		case v := <-ch1:
			fmt.Printf("v1 = %d\n", v)
		case v := <-ch2:
			fmt.Printf("v2 = %d\n", v)
		default:
			fmt.Println("get timeout.")
			time.Sleep(time.Second)
		}
	}
}
