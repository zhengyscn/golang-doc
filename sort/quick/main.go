package main

import (
	"fmt"
	"math/rand"
)

func quickSort(s []int) []int {
	/*
	   先定一个元素， 左边比这个元素小，右边比这个元素大。
	*/
	if len(s) <= 1 {
		return s
	}
	var left []int
	var right []int
	var retdata []int
	var first_value int = s[0]
	for _, v := range s[1:] {
		if v > first_value {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	leftSlice := quickSort(left)
	rightSlice := quickSort(right)
	retdata = append(retdata, leftSlice...)
	retdata = append(retdata, first_value)
	retdata = append(retdata, rightSlice...)
	return retdata

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
	s1 := quickSort(s)
	fmt.Println(s1)
}
