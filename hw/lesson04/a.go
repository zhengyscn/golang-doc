package main

import (
	"fmt"
	"sort"
)

func main() {
	var ids = []int{2, 5, 1, 7}
	sort.Ints(ids)
	fmt.Println(ids)
	fmt.Println(ids[len(ids)-1])
}
