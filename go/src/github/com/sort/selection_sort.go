package sort

func SelectionSort(data []int) {
	size := len(data)
	for i := 0; i < size - 1; i++ {
		idx := i
		for j := i + 1; j < size; j++ {
			if data[j] < data[idx] {
				idx = j
			}	
		}

		if (idx != i) {
			data[idx], data[i] = data[i], data[idx]
		}
	}
}