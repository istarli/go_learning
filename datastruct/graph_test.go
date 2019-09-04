package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultGraph() [][]int {
	/*
			 (10)
			0----1
		 (5)|	 |(5)
			3----2
			  (5)
	*/
	return [][]int{
		{0, 10, 0, 5},
		{10, 0, 5, 0},
		{0, 5, 0, 5},
		{5, 0, 5, 0},
	}
}

func targetPrimGraph() [][]int {
	/*
			0    1
		 (5)|	 |(5)
			3----2
			  (5)
	*/
	return [][]int{
		{0, 0, 0, 5},
		{0, 0, 5, 0},
		{0, 5, 0, 5},
		{5, 0, 5, 0},
	}
}

func targetDijstraGraph() [][]int {
	/*
			 (10)
			0----1
		 (5)|
			3----2
			  (5)
	*/
	return [][]int{
		{0, 10, 0, 5},
		{10, 0, 0, 0},
		{0, 0, 0, 5},
		{5, 0, 5, 0},
	}
}

func TestDFS(t *testing.T) {
	expectOrder := []int{0, 1, 2, 3}
	visitOrder := make([]int, 0)
	mat := defaultGraph()
	visited := make([]bool, len(mat))
	err := DFS(0, mat, visited, func(k int) {
		visitOrder = append(visitOrder, k)
	})
	assert.Nil(t, err)
	assert.True(t, EqualIntSlice(expectOrder, visitOrder))
}

func TestBFS(t *testing.T) {
	expectOrder := []int{0, 1, 3, 2}
	visitOrder := make([]int, 0)
	mat := defaultGraph()
	err := BFS(0, mat, func(k int) {
		visitOrder = append(visitOrder, k)
	})
	assert.Nil(t, err)
	assert.True(t, EqualIntSlice(expectOrder, visitOrder))
}

func TestPrim(t *testing.T) {
	mat := defaultGraph()
	res, err := Prim(0, mat)
	assert.Nil(t, err)
	assert.True(t, EqualIntMatrix(res, targetPrimGraph()))
}

func TestDijstra(t *testing.T) {
	mat := defaultGraph()
	res, err := Dijstra(0, mat)
	assert.Nil(t, err)
	assert.True(t, EqualIntMatrix(res, targetDijstraGraph()))
}
