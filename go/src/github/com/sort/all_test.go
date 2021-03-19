package sort

import (
	"math/rand"
	"reflect"
	"testing"
	_ "time"
)

var size = 10000

var tests = []struct{
	data     []int
	expected []int
} {
	{[]int {1},       []int {1}},
	{[]int {1, 2},    []int {1, 2}},
	{[]int {2, 1},    []int {1, 2}},
	{[]int {1, 2, 3}, []int{1, 2, 3}},
	{[]int {1, 3, 2}, []int{1, 2, 3}},
	{[]int {2, 1, 3}, []int{1, 2, 3}},
	{[]int {2, 3, 1}, []int{1, 2, 3}},
	{[]int {3, 2, 1}, []int{1, 2, 3}},
	{[]int {3, 1, 2}, []int{1, 2, 3}},
	{[]int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

type SortFunc func([]int)

func isSorted(data []int) bool {
	size := len(data)
	for i := 0; i < size - 1; i++ {
		if data[i] > data[i + 1] {
			return false
		}
	}

	return true
}

func sortNormal(t *testing.T, sf SortFunc) {
	for _, test := range tests {
		if sf(test.data); !reflect.DeepEqual(test.data, test.expected) {
			t.Errorf("expected: %v, got: %v\n", test.expected, test.data)
		}
	}
}

func sortRandom(t *testing.T, sf SortFunc) {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}

	sf(data)
	if !isSorted(data) {
		t.Errorf("expected: true, got: false\n")
	}
}

func TestSelectionSort(t *testing.T) {
	sortNormal(t, SelectionSort)
}

func TestSelectionSortRandom(t *testing.T) {
	sortRandom(t, SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	sortNormal(t, BubbleSort)
}

func TestBubbleSortRandom(t *testing.T) {
	sortRandom(t, BubbleSort)
}

func TestInsertionSort(t *testing.T) {
	sortNormal(t, InsertionSort)
}

func TestInsertionSortRandom(t *testing.T) {
	sortRandom(t, InsertionSort)
}

func TestMergeSort(t *testing.T) {
	sortNormal(t, MergeSort)
}

func TestMergeSortRandom(t *testing.T) {
	sortRandom(t, MergeSort)
}

func TestQuickSort(t *testing.T) {
	sortNormal(t, QuickSort)
}

func TestQuickSortRandom(t *testing.T) {
	sortRandom(t, QuickSort)
}

func TestHeapSort(t *testing.T) {
	sortNormal(t, HeapSort)
}

func TestHeapSortRandom(t *testing.T) {
	sortRandom(t, HeapSort)
}

func benchmark(b *testing.B, sf SortFunc, size int) {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int() % size
	}
	sf(data)
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmark(b, BubbleSort, size)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmark(b, InsertionSort, size)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmark(b, SelectionSort, size)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmark(b, MergeSort, size)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmark(b, QuickSort, size)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmark(b, HeapSort, size)
}

