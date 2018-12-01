package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	WithTimeout
	WithValue
	goroutine 优雅退出
*/
func worker(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	// interface to string ->
	trace_code, ok := ctx.Value("TRACE_CODE").(string)
	if !ok {
		fmt.Println("get trace_cod err")
	}

LOOP:
	for {
		fmt.Printf("hello world, trace_code %v\n", trace_code)
		time.Sleep(time.Second * 1)
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
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	ctx = context.WithValue(ctx, "TRACE_CODE", "456312345679")
	wg.Add(1)
	go worker(ctx, wg)
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("main done.")
}
