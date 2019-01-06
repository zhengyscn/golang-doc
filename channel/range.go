package main

import "fmt"

func main()  {
    var pipeChan = make(chan int, 10)
    for i:=0; i<10; i++ {
        pipeChan <- i
    }
    close(pipeChan) // 关闭channel

    for v := range pipeChan {
        fmt.Printf("value :%d\n", v)
    }

    fmt.Println("channel close end.")
}
