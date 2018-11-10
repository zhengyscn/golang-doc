package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Age      int
	Sex      string
}

/*
	reflect: 反射
*/

func updateInterfaceValue(a interface{}) { // 必须传指针类型才能修改值
	// 1. 获取a的类型
	// 2. 动态改变a里面的值
	// 3. 如果a存储的是一个结构体，还可以通过反射调用结构体里中的字段信息以及调用结构体里面的方法

	v := reflect.ValueOf(a)
	t := v.Type()
	switch t.Kind() {
	case reflect.Ptr:
		t1 := v.Elem().Type() // 指针指向变量的类型
		switch t1.Kind() {
		case reflect.Int:
			fmt.Println("Int")
			v.Elem().SetInt(200)
		case reflect.String:
			v.Elem().SetString("nihao")
		case reflect.Struct:
			fieldNum := t.NumField()
			fmt.Printf("fieldNum %d\n", fieldNum)
			//for i := 0; i < fieldNum; i++ {
			//	fmt.Println(t.Field(i))
			//}
		}
	}
}

func main() {
	var x = 100
	var y = "dajiahao"
	var u User = User{
		Username: "zhengyscn",
		Age:      28,
		Sex:      "male",
	}
	fmt.Printf("x=%d, y=%v, u=%v\n", x, y, u)
	updateInterfaceValue(&x)
	updateInterfaceValue(&y)
	updateInterfaceValue(&u)
	fmt.Printf("x=%d, y=%v, u=%v\n", x, y, u)
}
