package main

import (
    "fmt"
    "math/rand"
    "sort"
)

type Student struct {
    Id int
    Name string
    Age int
}

type Book struct {
    Id int
    Name string
    Author string
}

// 数组默认传地址
type stusArr []Student

func (this stusArr) Len() int {
    return len(this)
}

func (this stusArr) Less(i, j int) bool  {
    return this[i].Name < this[j].Name
}

func (this stusArr) Swap(i, j int)  {
    this[i], this[j] = this[j], this[i]
}

func main()  {
    var stus stusArr
    for i:=0; i<10; i++ {
        stu := Student{
            Name:fmt.Sprintf("zhengyscn-%d", rand.Intn(100)),
            Id:i,
            Age:rand.Intn(100),
        }
        stus = append(stus, stu)
    }
    fmt.Println("gen []student")
    for _, v := range stus {
        fmt.Printf("%#v\n", v)
    }

    fmt.Println("after sort")

    // 实现Sort的所有方法就实现了Sort接口
    sort.Sort(stus)
    for _, v := range stus {
        fmt.Printf("%#v\n", v)
    }
}
