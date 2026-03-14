package clrs

import (
	"cmp"
	"testing"
	"fmt"
)

const SmallBenchSize = 100
const BigBenchSize = 100000

type InPlaceSortFunc[T cmp.Ordered] func([]T)

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
	{9, 1, 2, 8, 7, 3, 4, 6, 5, 0},
	{5, 5, 6, 6, 4, 4, 7, 7, 8, 8, 3, 3, 2, 2, 1, 1, 9, 9, 0, 0},
}

func TestInsertionSort(t *testing.T) {
	for _, test := range tests {
		n := len(test)
		work := make([]int, n)
		copy(work, test)
		InsertionSort(work)
		if !IsSorted(work) {
			t.Errorf("InsertionSort(%v) sorted as %v\n", test, work)
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
	for _, test := range tests {
		n := len(test)
		work := make([]int, n)
		copy(work, test)
		SelectionSort(work)
		if !IsSorted(work) {
			t.Errorf("SelectionSort(%v) sorted as %v\n", test, work)
		}
	}
}

func BenchmarkSelectionSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, SmallBenchSize)
}

func BenchmarkSelectionSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, SelectionSort, BigBenchSize)
}

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		n := len(test)
		work := make([]int, n)
		copy(work, test)
		MergeSort(work)
		fmt.Printf("original: %v\n", test);
		fmt.Printf("sorted:   %v\n", work);
		if !IsSorted(work) {
			t.Errorf("SelectionSort(%v) sorted as %v\n", test, work)
		}
	}
}

func TestMergeSortDebug(t *testing.T) {
	orig := []int{4, 1, 2, 3}
	n := len(orig)
	work := make([]int, n)
	copy(work, orig)
	MergeSort(work)
	fmt.Println("sorted", orig, "as", work)
}

func BenchmarkMergeSortSmall(b *testing.B) {
	benchmarkInPlaceSort(b, MergeSort, SmallBenchSize)
}

func BenchmarkMergeSortBig(b *testing.B) {
	benchmarkInPlaceSort(b, MergeSort, BigBenchSize)
}

func benchmarkInPlaceSort(b *testing.B, f InPlaceSortFunc[int], size int) {
	for b.Loop() {
		b.StopTimer()
		items := RandInts(0, size, size)
		b.StartTimer()
		f(items)
	}
}
