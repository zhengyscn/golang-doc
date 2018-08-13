package test

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		items := make([]int, 10)
		for j := 0; j < 50; j++ {
			items = append(items, rand.Intn(100))
		}
		t.Log("before: ", items)
		BubbleSort(items)
		t.Log("after: ", items)
	}
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		items := make([]int, 10)
		for j := 0; j < 50; j++ {
			items = append(items, rand.Intn(100))
		}
		t.Log("before: ", items)
		QuickSort(items)
		t.Log("after: ", items)
	}
}
