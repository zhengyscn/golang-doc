package test

// 冒泡排序
func BubbleSort(items []int) {
	var (
		n       = len(items)
		swapret = true
	)
	for swapret {
		swapret = false
		for i := 0; i < n-1; i++ {

			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapret = true
			}
		}
		n = n - 1
	}
}

// 快排
func QuickSort(items []int) {
	sort(items, 0, len(items)-1)
}

func sort(items []int, left int, right int) {
	if left < right {
		var i, j = left, right
		var key = items[i]
		for i < j {
			for i < j && items[j] >= key {
				j--
			}
			items[i] = items[j]
			for i < j && items[i] <= key {
				i++
			}
			items[j] = items[i]
		}
		items[i] = key
		sort(items, left, i-1)
		sort(items, i+1, right)
	}
}
