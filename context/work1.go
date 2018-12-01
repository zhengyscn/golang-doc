package main

import (
	"fmt"
	"time"
)

func worker() {
	for {
		fmt.Println("hello world.")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go worker()
	time.Sleep(time.Second * 3)
	fmt.Println("main done.")
}
