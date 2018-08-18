package test

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel() //并行
	for i := 0; i < 10; i++ {
		items := make([]int, 10)
		for j := 0; j < 50; j++ {
			items = append(items, rand.Intn(100))
		}
		t.Log(t.Name(), "before: ", items)
		BubbleSort(items)
		t.Log(t.Name(), "after: ", items)
	}
}

func TestQuickSort(t *testing.T) {
	t.Parallel() //并行
	for i := 0; i < 10; i++ {
		items := make([]int, 10)
		for j := 0; j < 50; j++ {
			items = append(items, rand.Intn(100))
		}
		t.Log(t.Name(), "before: ", items)
		QuickSort(items)
		t.Log(t.Name(), "after: ", items)
	}
}
