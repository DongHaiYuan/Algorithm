package sort

func InsertionSort(data []int) {
	size := len(data)
	for i := 1; i < size; i++ {
		j, key := i - 1, data[i]
		for ; j >= 0; j-- {
			if key < data[j] {
				data[j + 1] = data[j]
			} else {
				break
			}
		}
		data[j + 1] = key
	}
}