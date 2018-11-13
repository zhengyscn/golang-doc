package main

import "fmt"

func console() {
	/*
		标准输出
	*/
	fmt.Print("hello world")
	fmt.Printf("hello world, %d\n", 100)
	fmt.Println("hello world")

	s1 := fmt.Sprint("hello world")
	fmt.Println(s1)

	s2 := fmt.Sprintf("hello world, %d", 100)
	fmt.Println(s2)

	var name string
	/*
		标准输入
		阻塞模式
	*/
	fmt.Scan(&name)
	fmt.Println(name)
}

func china() {
	var c rune
	c = '中'
	fmt.Printf("c=%c\n", c)
	fmt.Printf("c=%d\n", c)

	fmt.Printf("c=%c\n", 20013)
}

func main() {
	console()

	china()
}
