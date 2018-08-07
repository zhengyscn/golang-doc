package test

// BubbleSort xxx
func BubbleSort(items []int) {
	var (
		length = len(items)
		//swapret = true
	)
	for i := 0; i < length; i++ {
		if items[i] < items[i+1] {
			items[i], items[i+1] = items[i+1], items[i]
		}
	}
}
