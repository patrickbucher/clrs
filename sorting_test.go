package clrs

import "testing"

const SmallBenchSize = 100
const BigBenchSize = 100000

func TestInsertionSort(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2},
		{2, 1},
		{1, 2, 3},
		{2, 1, 3},
		{3, 2, 1},
		{3, 1, 2},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{5, 6, 4, 7, 3, 8, 2, 9, 1, 0},
		{0, 5, 6, 4, 7, 3, 8, 2, 9, 1},
		{-45, 45, 99},
		{99, 45, -45},
	}
	for _, test := range tests {
		n := len(test)
		original := make([]int, n)
		copy(original, test)
		InsertionSort(test)
		if !IsSorted(test) {
			t.Errorf("InsertionSort(%v) sorted as %v\n", original, test)
		}
	}
}

func BenchmarkInsertionSortSmall(b *testing.B) {
	benchmarkInsertionSort(b, SmallBenchSize)
}

func BenchmarkInsertionSortBig(b *testing.B) {
	benchmarkInsertionSort(b, BigBenchSize)
}

func benchmarkInsertionSort(b *testing.B, size int) {
	for b.Loop() {
		b.StopTimer()
		items := RandInts(0, size, size)
		b.StartTimer()
		InsertionSort(items)
	}
}
