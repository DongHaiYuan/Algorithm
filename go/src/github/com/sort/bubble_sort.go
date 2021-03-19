package sort

func BubbleSort(data []int) {
	size := len(data)
	for i := 0; i < size - 1; i++ {
		flag := false
		for j := 0; j < size - i - 1; j++ {
			if data[j] > data[j + 1] {
				flag = true
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
		if !flag {
			break
		}
	}
}