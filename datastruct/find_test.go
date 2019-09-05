package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	seq := []int{-1, 3, 5, 7, 8}
	assert.Equal(t, 0, BinarySearch(seq, -1))
	assert.Equal(t, 1, BinarySearch(seq, 3))
	assert.Equal(t, len(seq)-1, BinarySearch(seq, 8))
	assert.Equal(t, -1, BinarySearch(seq, -2))
	assert.Equal(t, -1, BinarySearch(seq, 9))
}

func TestLowerBound(t *testing.T) {
	seq := []int{-1, 3, 5, 7, 8}
	assert.Equal(t, 0, LowerBound(seq, -2))
	assert.Equal(t, 0, LowerBound(seq, -1))
	assert.Equal(t, 1, LowerBound(seq, 3))
	assert.Equal(t, 2, LowerBound(seq, 4))
	assert.Equal(t, len(seq)-1, LowerBound(seq, 8))
	assert.Equal(t, len(seq), LowerBound(seq, 9))
}

func TestUpperBound(t *testing.T) {
	seq := []int{-1, 3, 5, 7, 8}
	assert.Equal(t, 0, UpperBound(seq, -2))
	assert.Equal(t, 1, UpperBound(seq, -1))
	assert.Equal(t, 2, UpperBound(seq, 3))
	assert.Equal(t, 2, UpperBound(seq, 4))
	assert.Equal(t, len(seq), UpperBound(seq, 8))
	assert.Equal(t, len(seq), UpperBound(seq, 9))
}
