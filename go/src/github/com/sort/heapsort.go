package sort

// 1 2 3 4 5
// heapify => 3 2 1
// 
func HeapSort(data []int) {
	heapify(data)
	size := len(data)
	for i := 0; i < size - 1; i++ {
		last := size - i - 1
		exch(data, 0, last)
		sink(data, 0, last)
	}
}

func heapify(data []int) {
	size := len(data)
	for i := size - 1; i > 0; i-- {
		p := parent(i)
		if data[i] > data[p] {
			exch(data, i, p)
			sink(data, i, size)
		}
	}
}

func parent(index int) int {
	if index % 2 == 0 {
		return index / 2 - 1
	}

	return index / 2
}

func sink(data []int, index int, size int) {
	for index < size {
		tmp := index
		left := index * 2 + 1
		right := index * 2 + 2
		if left < size && data[tmp] < data[left] {
			tmp = left
		}
		if right < size && data[tmp] < data[right] {
			tmp = right
		}
		if tmp == index {
			break
		}
		exch(data, tmp, index)
		index = tmp
	}
}

// exchange data in the position at i and j
func exch(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}