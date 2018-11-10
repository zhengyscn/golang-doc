package main

import "fmt"

/*
	& 取指针的地址
	* 取指针的地址的值

	%T	表示变量类型
	%p  表示指针地址
	%v	表示任意类型
*/

func example_v1() {
	var i int32
	fmt.Printf("value of i %d\n", i)
	i = 100
	fmt.Printf("value of i %d\n", i)
	fmt.Printf("address of i %v\n", &i)
	fmt.Printf("address of i %p\n", &i)
	fmt.Printf("value of i %v\n", *(&i))
}

func example_v2() {
	var i *int32
	var j int32 = 100
	i = &j
	fmt.Printf("value of i %d\n", *i)
	fmt.Printf("address of i %p\n", i)
}

func example_v3() {
	var i *int
	var j int = 100
	i = &j
	fmt.Printf("value of i %d\n", i)
	fmt.Printf("address of i %p\n", &i)
	fmt.Printf("address of i %p\n", &j)
	*i++
	fmt.Println(j)
}

func example_v4() {
	/*
		nil
		空地址不允许赋值
	*/
	var a *int
	var b int = 12
	*a = b
	fmt.Println(a)
	//if a != nil {
	//	*a = 100
	//	fmt.Printf("value of a %d", a)
	//	fmt.Printf("address of a %d", a)
	//} else {
	//	fmt.Printf("value of a %d\n", a)
	//}

}

func modify_value() {
	var b int = 255
	a := &b
	fmt.Printf("address of a %p\n", a)
	fmt.Printf("value of a %v\n", *a)
	*a++
	fmt.Printf("value of a %v\n", *a)
}

func modify_ptr_v1(x int) {
	x++
}

func modify_ptr_v2(x *int) {
	*x++
}

func main() {
	//example_v1()
	//example_v2()
	//example_v3()
	//example_v4()
	//modify_value()
	var x1 int = 2
	modify_ptr_v1(x1)
	fmt.Println(x1)

	var x2 int = 2
	modify_ptr_v2(&x2)
	fmt.Println(x2)
}
