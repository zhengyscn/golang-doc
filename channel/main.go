package main

import (
    "fmt"
)

/*

    channel 信道
    1. 无缓冲区的channel不负责存储数据，只负责数据的流通；
    2. 无缓冲区的channel顺序是先进先出；

*/

func pipe_chan1()  {
    /*
        无缓冲区的channel
     */
    ch := make(chan int)

    go func() {
        ch <- 100
    }()
    fmt.Println(<-ch) // 堵塞

    fmt.Println("Pipe Chan Over.")
}

func pipe_chan2()  {
    /*
        无缓冲区的channel
     */
    ch := make(chan int)
    go func() {
        fmt.Println(<-ch)
    }()
    ch <- 1
    fmt.Println("Pipe Chan Over.")
}

func pipe_chan3()  {
    /*
        无带缓冲区的channel
     */
    ch := make(chan int)
    go func() {
        for pipe := range ch {
            fmt.Printf("pipe = %d\n", pipe)
        }
    }()

    for i:=0; i<5; i++ {
        ch <- i
    }
    fmt.Println("Pipe Chan Over.")
}

func pipe_buf_chan1()  {
    /*
        无带缓冲区的channel
     */

    pipeCh := make(chan int, 2)
    pipeCh <- 1
    pipeCh <- 2
    fmt.Println(<-pipeCh)
    fmt.Println(<-pipeCh)
    fmt.Println("Pipe Chan Over.")
}

func pipe_buf_chan2()  {
    /*
        无带缓冲区的channel
     */

    pipeCh := make(chan int, 3)
    pipeCh <- 1
    pipeCh <- 2
    pipeCh <- 3
    for ch := range pipeCh {
        fmt.Println(ch)
        if len(pipeCh) == 0 {
            break
        }
    }
    fmt.Println("Pipe Chan Over.")
}


func pipe_buf_chan3()  {
    /*
        无带缓冲区的channel
     */

    pipeCh := make(chan int, 3)

    go func() {
        for ch := range pipeCh {
            fmt.Println(ch)
            if len(pipeCh) == 0 {
                break
            }
        }
    }()

    for i:=0; i<5; i++ {
        pipeCh <- i
    }
    fmt.Println("Pipe Chan Over.")
}


func main()  {
    //pipe_chan1()
    //pipe_chan2()
    //pipe_chan3()
    //pipe_buf_chan1()
    //pipe_buf_chan2()
    pipe_buf_chan3()
}