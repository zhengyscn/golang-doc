package test

import "testing"

func TestBubbleSort(t *testing.T) {
	items := []int{2, 1, 5, 3}
	BubbleSort(items)

	t.Logf("Before sort items, %s", items)
	t.Logf("After sort items, %s", items)
}
