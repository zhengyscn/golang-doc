package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

/*
    终端读写
    文件读写

    os.Stdin    标准输入
    os.Stdout   标准输出
    os.Stderr   标准错误输出

    https://studygolang.com/articles/11610?fr=sidebar
 */

func test1()  {
    fd, err := os.OpenFile("terminal.log", os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("open file err:%v\n", err)
        return
    }

    for {
        io.Copy(fd, os.Stdin)
    }
}

type User struct {
    Username string
    Age int
    Tel string
}

func can1()  {
    var username, age, tel string
    fmt.Printf("Please input user info: ")
    // Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
    fmt.Scanln(&username, &age, &tel)
    fmt.Printf("username:%s, age:%s, tel:%s\n", username, age, tel)
}

func can2()  {
    var username, age, tel string
    fmt.Printf("Please input user info: ")
    // Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
    fmt.Scanf("%s %s %s", &username, &age, &tel)
    fmt.Printf("username:%s, age:%s, tel:%s\n", username, age, tel)
}

func can3()  {
    var username, tel string
    var age int
    // Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，
    // 与 Scanf 相同。如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误

    fmt.Sscanf("user01 / 18 / 10010", "%s / %d / %s", &username, &age, &tel)
    fmt.Printf("username:%s, age:%d, tel:%s\n", username, age, tel)
}

// 带缓冲区的读写
func input1()  {
    // inputReader 是一个指向 bufio.Reader 的指针。
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Printf("read string \n")
        return
    }
    fmt.Println(input)
}

func input2()  {
    // inputReader 是一个指向 bufio.Reader 的指针。
    var username, tel string
    var age int
    for {
        fmt.Printf("Please input user info: ")
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("read string \n")
            return
        }
        _, err = fmt.Sscan(input, &username, &age, &tel)
        if err != nil {
            fmt.Printf("sscan err :%v\n", err)
            return
        }
        fmt.Println(username, age, tel)
    }

}

func main()  {
    input2()
}
