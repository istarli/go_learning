package data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultGraph() [][]int {
	res := [][]int{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
	}
	return res
}
func TestDFS(t *testing.T) {
	expectOrder := []int{0, 1, 2, 3}
	visitOrder := make([]int, 0)
	mat := defaultGraph()
	visited := make([]bool, len(mat))
	DFS(0, mat, visited, func(k int) {
		visitOrder = append(visitOrder, k)
	})
	assert.True(t, EqualIntSlice(expectOrder, visitOrder))
}
