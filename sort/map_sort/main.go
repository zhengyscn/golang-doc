package main

import (
	"fmt"
	"sort"
)

/*
   map 是无序的， 每一次打印有可能就会和上一次不一样。
*/

func main() {
	var keys []int
	var m map[int]string = make(map[int]string, 5)

	m[6] = "zhengyansheng"
	m[10] = "zhengyansheng"
	m[15] = "zhengyansheng"
	m[4] = "zhengyansheng"
	m[1] = "zhengyansheng"

	fmt.Println(">>> before sort.Ints")
	for k, v := range m {
		keys = append(keys, k)
		fmt.Printf("[%d %s]\n", k, v)
	}

	sort.Ints(keys)

	fmt.Println()
	fmt.Println(">>> after sort.Ints")

	for _, k := range keys {
		fmt.Printf("[%d %s]\n", k, m[k])
	}

}
