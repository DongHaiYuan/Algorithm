package sort

// 2 1   -> r = -1, l = 0
// 3 1 2 -> r = 0, l = 1
// 5 2 1 3 -> r = 
func partition(data []int, left, right int) int {
	pivot := data[right]
	r := right - 1
	l := left
	for ; r >= l; {
		for ; r >= l && data[r] > pivot; r-- {}
		for ; r >= l && data[l] < pivot; l++ {}
		if r >= l {
			data[r], data[l] = data[l], data[r]
			r--
			l++
		} 
	}

	data[l], data[right] = data[right], data[l]

	return l
}

func quickSort(data []int, left, right int) {
	if left < right {
		pv := partition(data, left, right)
		quickSort(data, left, pv - 1)
		quickSort(data, pv + 1, right)
	}
}

func QuickSort(data []int) {
	quickSort(data, 0, len(data) - 1)
}