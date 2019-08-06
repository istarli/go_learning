package data_structure

func DFS(k int, mat [][]int, visited []bool, fn func(int)) {
	if k < 0 || k > len(mat) {
		return
	}
	visited[k] = true
	fn(k)
	for i, weight := range mat[k] {
		if !visited[i] && linkOn(weight) {
			DFS(i, mat, visited, fn)
		}
	}
}

func linkOn(weight int) bool {
	return weight > 0
}
