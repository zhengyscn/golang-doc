package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Nickname string
	Age      int
	Tel      string
}

func mapUnmarshal() {
	var u User
	var user01 = `{"Username":"user01","Nickname":"用户1","Age":100,"Tel":"10010"}`

	/*
	   Unmarshal
	   - data []byte     : json字符串
	   - v interface{}   : 传指针
	*/
	err := json.Unmarshal([]byte(user01), &u)
	if err != nil {
		fmt.Printf("unmarshal map err:%v\n", err)
		return
	}
	fmt.Println(reflect.TypeOf(u), u)
}

func main() {
	mapUnmarshal()
}
