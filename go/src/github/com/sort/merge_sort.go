package sort

func merge(left []int, right []int) []int {
	size1 := len(left)
	size2 := len(right)
	mg := make([]int, size1 + size2)
	i, j, k := 0, 0, 0
	for ; i < size1 && j < size2; {
		if left[i] < right[j] {
			mg[k] = left[i]
			i++
		} else {
			mg[k] = right[j]
			j++
		}
		k++
	}

	for ; i < size1; i++ {
		mg[k] = left[i]
		k++
	}

	for ; j < size2; j++ {
		mg[k] = right[j]
		k++
	}

	return mg
}

func mergeSort(data []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(data, left, mid)
		mergeSort(data, mid + 1, right)
		tmp := merge(data[left : mid + 1], data[mid + 1 : right + 1])
		for i := 0; i < len(tmp); i++ {
			data[left + i] = tmp[i]
		}
	}
}

func MergeSort(data []int) {
	mergeSort(data, 0, len(data) - 1)
}