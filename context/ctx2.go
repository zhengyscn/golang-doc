package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	WithTimeout
	goroutine 优雅退出
*/
func worker(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Microsecond * 1)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	wg.Add(1)
	go worker(ctx, wg)
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("main done.")
}
