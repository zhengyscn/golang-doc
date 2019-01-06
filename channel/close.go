package main

import "fmt"

func main()  {
    var pipeChan = make(chan int, 10)
    for i:=0; i<10; i++ {
        pipeChan <- i
    }
    close(pipeChan) // 关闭channel
    for {
        value, ok := <- pipeChan
        if !ok {
            fmt.Printf("pipeChan %d\n", value)
            break
        }
        fmt.Printf("value :%d\n", value)
    }
    fmt.Println("channel close end.")
}
