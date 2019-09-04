package datastruct

import (
	"fmt"
	"math"
)

const (
	maxWeight = math.MaxInt32
)

//  <= 0 is block, > 0 is weight
func linkOn(weight int) bool {
	return weight > 0
}

func toCost(weight int) int {
	if linkOn(weight) {
		return weight
	}
	return maxWeight
}

func toPre(k, weight int) int {
	if linkOn(weight) {
		return k
	}
	return -1
}

// DFS ...
func DFS(k int, mat [][]int, visited []bool, fn func(int)) error {
	if len(mat) > 0 && len(mat) != len(mat[0]) {
		return fmt.Errorf("Invalid mat[][]")
	}
	if k < 0 || k >= len(mat) {
		return fmt.Errorf("Invalid k")
	}
	if len(visited) < len(mat) {
		return fmt.Errorf("Invalid visited[]")
	}
	if fn == nil {
		return fmt.Errorf("Invalid fn")
	}
	dfsCore(k, mat, visited, fn)
	return nil
}

func dfsCore(k int, mat [][]int, visited []bool, fn func(int)) {
	fn(k)
	visited[k] = true
	for i, weight := range mat[k] {
		if !visited[i] && linkOn(weight) {
			dfsCore(i, mat, visited, fn)
		}
	}
}

// BFS ...
func BFS(k int, mat [][]int, fn func(int)) error {
	if len(mat) > 0 && len(mat) != len(mat[0]) {
		return fmt.Errorf("Invalid mat[][]")
	}
	if k < 0 || k >= len(mat) {
		return fmt.Errorf("Invalid k")
	}
	if fn == nil {
		return fmt.Errorf("Invalid fn")
	}
	bfsCore(k, mat, fn)
	return nil
}

func bfsCore(k int, mat [][]int, fn func(int)) {
	q := NewQueue()
	visited := make([]bool, len(mat))
	q.Push(k)
	visited[k] = true
	for !q.Empty() {
		v, _ := q.Pop()
		fn(v)
		for i, weight := range mat[v] {
			if !visited[i] && linkOn(weight) {
				q.Push(i)
				visited[i] = true
			}
		}
	}
}

// Prim ...
func Prim(k int, mat [][]int) ([][]int, error) {
	if len(mat) > 0 && len(mat) != len(mat[0]) {
		return nil, fmt.Errorf("Invalid mat[][]")
	}
	if k < 0 || k >= len(mat) {
		return nil, fmt.Errorf("Invalid k")
	}
	return spanningTreeCore(k, mat, func(cost, deta int) int { return deta }), nil
}

// Dijstra ...
func Dijstra(k int, mat [][]int) ([][]int, error) {
	if len(mat) > 0 && len(mat) != len(mat[0]) {
		return nil, fmt.Errorf("Invalid mat[][]")
	}
	if k < 0 || k >= len(mat) {
		return nil, fmt.Errorf("Invalid k")
	}
	return spanningTreeCore(k, mat, func(cost, deta int) int {
		if cost == maxWeight || deta == maxWeight {
			return maxWeight
		}
		return cost + deta
	}), nil
}

func spanningTreeCore(k int, mat [][]int, fn func(int, int) int) [][]int {
	n := len(mat)
	// init help-data
	visited := make([]bool, n)
	visited[k] = true
	cost := make([]int, n)
	pre := make([]int, n)
	for i, w := range mat[k] {
		cost[i] = toCost(w)
		pre[i] = toPre(k, w)
	}
	// n-1 is enough, i has no meaning inside for-body
	for i := 0; i < n-1; i++ {
		// Step1, find min cost
		v, min := k, maxWeight
		for j := 0; j < n; j++ {
			if !visited[j] && cost[j] < min {
				min = cost[j] // Caution, can not use cost[v] to replace min here.
				v = j
			}
		}
		// Step2, if no valid node, break
		if v == k {
			break
		}
		// Step3, update help-data
		visited[v] = true
		for j := 0; j < n; j++ {
			newCost := fn(cost[v], toCost(mat[v][j]))
			if !visited[j] && cost[j] > newCost {
				cost[j] = newCost
				pre[j] = v
			}
		}
	}
	// Generate res matrix by pre
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i, v := range pre {
		if v >= 0 && v < n {
			res[i][v], res[v][i] = mat[i][v], mat[v][i]
		}

	}
	return res
}
