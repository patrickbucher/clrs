package clrs

import (
	"cmp"
	"slices"
	"testing"
)

const SmallBenchSize = 100
const BigBenchSize = 100000

type InPlaceSortFunc[T cmp.Ordered] func([]T)

var tests = []struct {
	input    []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{2, 1, 3}, []int{1, 2, 3}},
	{[]int{4, 1, 2, 3}, []int{1, 2, 3, 4}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
}

func TestInsertionSort(t *testing.T) {
	testSort(t, InsertionSort)
}

func BenchmarkInsertionSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, InsertionSort, SmallBenchSize)
}

func BenchmarkInsertionSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, InsertionSort, BigBenchSize)
}

func TestSelectionSort(t *testing.T) {
	testSort(t, SelectionSort)
}

func BenchmarkSelectionSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, SmallBenchSize)
}

func BenchmarkSelectionSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, BigBenchSize)
}

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort)
}

func BenchmarkMergeSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, MergeSort, SmallBenchSize)
}

func BenchmarkMergeSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, MergeSort, BigBenchSize)
}

func testSort(t *testing.T, f InPlaceSortFunc[int]) {
	for _, test := range tests {
		n := len(test.input)
		work := make([]int, n)
		copy(work, test.input)
		f(work)
		if !slices.Equal(work, test.expected) {
			t.Errorf("sorted %v as %v\n", test, work)
		}
	}
}

func benchmarkInPlaceSort(b *testing.B, f InPlaceSortFunc[int], size int) {
	for b.Loop() {
		b.StopTimer()
		items := RandInts(0, size, size)
		b.StartTimer()
		f(items)
	}
}
