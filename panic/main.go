package main

import (
    "fmt"
    "time"
)

func test()  {
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("panic err:%v\n", err)
        }
    }()
    var m map[string]string
    m["name"] = "zhengyscn"
    fmt.Println(m)
}

func main()  {
    test()
    time.Sleep(time.Second * 3)
}
