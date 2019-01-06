package main

import (
	"fmt"
)

func send(ch chan int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 切记， 关闭channel，否则就deadlook
	fmt.Println(">>> send end.")
}

func recv(ch chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch // channel close, so ok is false.
		if !ok {
			break
		}
		fmt.Printf("v = %d\n", v)
	}
	exitChan <- struct{}{}
	fmt.Println(">>> recv end.")
}

func main() {
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 1)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	<-exitChan

}
