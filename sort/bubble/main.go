package main

import (
	"fmt"
	"math/rand"
)

// 冒泡排序
func bubbleSort(s []int) {
	// slice 引用类型
	for i := 0; i < len(s); i++ {
		// [89 87 47 59 81 18 25 40 56 0]
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func genSlice() (s []int) {
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(100))
	}
	return
}

func main() {
	s := genSlice()
	fmt.Println(s)
	bubbleSort(s)
	fmt.Println(s)
}
