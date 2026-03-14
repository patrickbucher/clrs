package clrs

import (
	"cmp"
	"testing"
)

const SmallBenchSize = 100
const BigBenchSize = 100000

type InPlaceSortFunc[T cmp.Ordered] func([]T)

func sortingTests() [][]int {
	var tests = [][]int{
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
	return tests
}

func TestInsertionSort(t *testing.T) {
	for _, test := range sortingTests() {
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
	benchmarkInPlaceSort(b, InsertionSort, SmallBenchSize)
}

func BenchmarkInsertionSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, InsertionSort, BigBenchSize)
}

func TestSelectionSort(t *testing.T) {
	for _, test := range sortingTests() {
		n := len(test)
		original := make([]int, n)
		copy(original, test)
		SelectionSort(test)
		if !IsSorted(test) {
			t.Errorf("SelectionSort(%v) sorted as %v\n", original, test)
		}
	}
}

func BenchmarkSelectionSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, SmallBenchSize)
}

func BenchmarkSelectionSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, BigBenchSize)
}

func benchmarkInsertionSort(b *testing.B, size int) {
	for b.Loop() {
		b.StopTimer()
		items := RandInts(0, size, size)
		b.StartTimer()
		InsertionSort(items)
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

func TestMergeSort(t *testing.T) {
	for _, test := range sortingTests() {
		n := len(test)
		original := make([]int, n)
		copy(original, test)
		MergeSort(test)
		if !IsSorted(test) {
			t.Errorf("SelectionSort(%v) sorted as %v\n", original, test)
		}
	}
}
