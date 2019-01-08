package main

import (
	"fmt"
	"time"
)

// 定义一个车接口
// 实现接口的所有方法，那么才实现了这个接口.
type Carer interface {
	// 方法签名 : 方法传参和返回值
	GetName() string
	Run()
	DiDi()
}

// 自行车
type Bicycle interface {
	// 车轮
	Wheel()
}

type BMW struct {
	Name string
}

func (this *BMW) GetName() string {
	return this.Name
}

func (this *BMW) Run() {
	for {
		fmt.Printf("%s is running.\n", this.Name)
		time.Sleep(time.Second * 1)
	}
}

func (this *BMW) DiDi() {
	fmt.Println("DiDi")
}

func (this *BMW) Wheel() {
	fmt.Println("wheel")
}

func main() {
	var car Carer
	var bicycle Bicycle
	var bmw = BMW{Name: "baoma"}
	car = &bmw
	car.DiDi()
	fmt.Println("------ car ------")
	bicycle = &bmw
	bicycle.Wheel()

}
