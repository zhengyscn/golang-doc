package main

import "fmt"

/*
	类型断言
	type assets
*/

func main() {
	var a int = 100
	var b string = "hello"
	var c interface{}

	c = a
	c1, ok := c.(int)
	if !ok {
		fmt.Println("c is not int")
	}
	fmt.Printf("c1=%d\n", c1)

	c = b
	c2, ok := c.(string)
	if !ok {
		fmt.Println("c is not string")
	}
	fmt.Printf("c2=%s\n", c2)

}
