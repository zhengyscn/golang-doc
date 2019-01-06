package main

import (
    "fmt"
    "reflect"
)

type student struct {
    id int
    name string
    age int
}

func chan1()  {
    var intChan chan int
    intChan = make(chan int, 1)
    intChan <- 10
}

func chan2()  {
    var stuChan chan interface{}
    stuChan = make(chan interface{}, 10)

    stu := student{
        id : 1,
        name:"xxx",
        age:29,
    }
    stuChan <- &stu

    var stu1 interface{}
    stu1 = <- stuChan
    fmt.Println(stu1)
    fmt.Println(reflect.TypeOf(stu1))

    /*
        类型断言
     */
    var stu2 *student
    stu2, ok := stu1.(*student)
    if !ok {
        fmt.Println("can not convert")
        return
    }
    fmt.Println(stu2)

}


func main()  {
    chan1()
    fmt.Println("chan 1 end")
    chan2()
    fmt.Println("chan 2 end")
}
