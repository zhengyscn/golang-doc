package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBubbleSort(b *testing.B) {
	items := []int{55, 3311, 99, 77}
	BubbleSort(items)
}

func BenchmarkQuickSort(b *testing.B) {
	items := []int{55, 3311, 99, 77}
	QuickSort(items)
}

func TestMain(m *testing.M) {
	fmt.Println("setup")
	rand.Seed(time.Now().Unix())
	m.Run()
	fmt.Println("tear down")
}
