package main

import (
	"fmt"
	"time"
)

var exitFlag bool

func worker() {
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Second * 1)
		if exitFlag {
			break
		}
	}
}

func main() {
	go worker()
	time.Sleep(time.Second * 3)
	exitFlag = true
	fmt.Println("main done.")
}
