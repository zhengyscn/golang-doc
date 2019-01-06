package main

import "fmt"

func readWrite()  {
    var intChan = make(chan int, 1)
    intChan <- 1
    fmt.Println(<-intChan)
}

func onlyRead()  {
    // 只读channel
    var ch1 <- chan int

    // 只写channel
    var ch2 chan <- int
}

func main()  {
    readWrite()
}
