package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultSeqs() ([]int, []int) {
	seq := []int{8, 3, 9, -1, 0, 4, 3}
	expect := []int{-1, 0, 3, 3, 4, 8, 9}
	return seq, expect
}

func assertEqualIntSlice(t *testing.T, seq []int, expect []int) {
	for i := range seq {
		assert.Equal(t, expect[i], seq[i])
	}
}

func assertEqualUintSlice(t *testing.T, seq []uint, expect []uint) {
	for i := range seq {
		assert.Equal(t, expect[i], seq[i])
	}
}

func TestInsertSort(t *testing.T) {
	seq, expect := defaultSeqs()
	InsertSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestShellSort(t *testing.T) {
	seq, expect := defaultSeqs()
	ShellSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestBubbleSort(t *testing.T) {
	seq, expect := defaultSeqs()
	BubbleSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestQuickSort(t *testing.T) {
	seq, expect := defaultSeqs()
	QuickSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestSelectSort(t *testing.T) {
	seq, expect := defaultSeqs()
	SelectSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestHeapSort(t *testing.T) {
	seq, expect := defaultSeqs()
	HeapSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestMergeSort(t *testing.T) {
	seq, expect := defaultSeqs()
	MergeSort(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestMergeSortNR(t *testing.T) {
	seq, expect := defaultSeqs()
	MergeSortNR(seq)
	assertEqualIntSlice(t, seq, expect)
}

func TestRadixSort(t *testing.T) {
	seq := []uint{3,1024,48,15,5894,93,23,66,5,31,5}
	expect := []uint{3,5,5,15,23,31,48,66,93,1024,5894}
	RadixSort(seq)
	t.Log(seq)
	assertEqualUintSlice(t,seq, expect)
}
