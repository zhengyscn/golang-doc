package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	stuinfo map[string]Student = make(map[string]Student, 10)
)

type Student struct {
	Id    string
	Name  string
	Age   int
	Sex   string
	Score float32
}

func getStuInfo() Student {
	var stu Student

	//fmt.Printf("please input id: ")
	//fmt.Scanf("%s\n", &stu.Id)

	fmt.Printf("please input name: ")
	fmt.Scanf("%s\n", &stu.Name)

	fmt.Printf("please input age: ")
	fmt.Scanf("%d\n", &stu.Age)

	fmt.Printf("please input sex: ")
	fmt.Scanf("%s\n", &stu.Sex)

	fmt.Printf("please input score: ")
	fmt.Scanf("%f\n", &stu.Score)

	return stu
}

func add() {
	stu := getStuInfo()

	if len(stuinfo) == 0 {
		stu.Id = "1"
		stuinfo["1"] = stu
	} else {
		fmt.Println("enter a message.")
		var (
			ids    []int
			max_id string
		)
		for k, _ := range stuinfo {
			i, _ := strconv.Atoi(k)
			ids = append(ids, i)
		}
		sort.Ints(ids)
		max_id = fmt.Sprintf("%d", ids[len(ids)-1]+1)
		stu.Id = max_id
		stuinfo[max_id] = stu
	}
}

func showMenu() {
	fmt.Println("please select:")
	fmt.Println("1. 添加用户信息")
	fmt.Println("2. 修改用户信息")
	fmt.Println("3. 显示用户信息")
	fmt.Println("4. 退出")
}

func main() {
	for {
		fmt.Printf("%#v\n", stuinfo)
		showMenu()
		var selectNum int
		fmt.Printf(">>> ")
		fmt.Scanf("%d\n", &selectNum)
		switch selectNum {
		case 1:
			add()
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("invalid select")
		}
	}
}
