package data_structure

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
