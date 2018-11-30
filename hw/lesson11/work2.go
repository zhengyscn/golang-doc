package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(exitChan chan struct{}, wg sync.WaitGroup) {
	defer wg.Done()
loop:
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Second * 1)
		switch {
		case <-exitChan:
			break loop
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
	go worker(exitChan, wg)
	wg.Wait()
	exitChan <- struct{}{}
	fmt.Println("main done.")
}
