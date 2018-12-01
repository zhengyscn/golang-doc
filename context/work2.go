package main

import (
	"fmt"
	"time"
)

/*
	goroutine 优雅退出
*/
func worker(exitFlag bool) {
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Second * 1)
		if exitFlag {
			break
		}
	}
}

func main() {
	var (
		exitFlag bool
	)
	go worker(exitFlag)
	time.Sleep(time.Second * 3)
	exitFlag = true
	fmt.Println("main done.")
}
