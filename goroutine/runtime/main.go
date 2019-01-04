package main

import (
    "fmt"
    "runtime"
)

func main()  {
    num := runtime.NumCPU()
    runtime.GOMAXPROCS(num)
    fmt.Printf("num cpu: %d\n", num)
}

