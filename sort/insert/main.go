package main

import (
	"fmt"
	"math/rand"
)

// 插入排序
func insertSort(s []int) {
	/*
	    slice 引用类型
	   [89 87 47 59 81 18 25 40 56 0]
	   第一个元素肯定是有序的
	*/
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
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
	insertSort(s)
	fmt.Println(s)
}
