package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	goroutine 优雅退出
*/
func worker(wg sync.WaitGroup, exitChan chan struct{}) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Second * 1)
		select {
		case <-exitChan:
			break LOOP
		default:

		}
	}
}

func main() {
	var (
		exitChan chan struct{} = make(chan struct{})
		wg       sync.WaitGroup
	)
	wg.Add(1)
	go worker(wg, exitChan)
	time.Sleep(time.Second * 5)
	exitChan <- struct{}{}
	fmt.Println("main done.")
}
