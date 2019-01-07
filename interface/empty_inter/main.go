package main

import (
	"fmt"
)

// 定义一个空接口
type A interface {
}

func aTest() {
	var a A
	var i int = 1
	var ii float64 = 1.5
	a = i
	fmt.Printf("a value %v, type %T\n", a, a)

	a = ii
	fmt.Printf("a value %v, type %T\n", a, a)
}

func main() {
	// empty interface
	aTest()
}
