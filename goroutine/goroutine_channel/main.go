package main

import (
	"fmt"
	"time"
)

/*
	goroutine如何通过channel通信
 */

func producter(ch chan int) {
	var i = 0
	for {
		if i == 5 {
			break
		}
		ch <- i
		i++
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer_v1(pipeCh chan int) {
	for {
		msg, ok := <- pipeCh
		if !ok {
			fmt.Printf("consumer pipe channel failed\n")
			return
		}
		fmt.Printf("%d < ch\n", msg)
	}
}


func consumer_v2(pipeCh chan int) {
	for ch := range pipeCh {
		fmt.Printf("%d < ch\n", ch)
	}
}

func consumer_v3(pipeCh chan int) {
	for {
		select {
		case msg := <-pipeCh:
			fmt.Printf("%d < ch\n", msg)
		}
	}
}

func main() {
	var pipeChan = make(chan int)
	go producter(pipeChan)
	consumer_v1(pipeChan)
	fmt.Println("main thread end.")
}
