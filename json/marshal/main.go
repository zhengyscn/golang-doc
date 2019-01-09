package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Nickname string
	Age      int
	Tel      string
}

func intMarshal() {
	var i int = 2
	bytes, err := json.Marshal(i)
	if err != nil {
		fmt.Printf("int marshal error: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

func mapMarshal() {
	var u User = User{
		Username: "user01",
		Nickname: "用户1",
		Age:      100,
		Tel:      "10010",
	}

	bytes, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("struct marshal error: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

func sliceMarshal() {
	var userSlice []User
	var u01 User = User{
		Username: "user01",
		Nickname: "用户1",
		Age:      100,
		Tel:      "10010",
	}
	userSlice = append(userSlice, u01)

	var u02 User = User{
		Username: "user01",
		Nickname: "用户1",
		Age:      100,
		Tel:      "10010",
	}
	userSlice = append(userSlice, u02)

	bytes, err := json.Marshal(userSlice)
	if err != nil {
		fmt.Printf("struct marshal error: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

func main() {
	// int marshal
	intMarshal()
	fmt.Println("-----1")

	// struct marshal
	mapMarshal()
	fmt.Println("-----2")

	// map marshal
	sliceMarshal()
	fmt.Println("-----3")
}
