package datastruct

// EqualIntSlice ...
func EqualIntSlice(seq1 []int, seq2 []int) bool {
	if len(seq1) != len(seq2) {
		return false
	}
	for i := range seq1 {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

// EqualIntMatrix ...
func EqualIntMatrix(mat1 [][]int, mat2 [][]int) bool {
	if len(mat1) != len(mat2) {
		return false
	}
	for i := range mat1 {
		if !EqualIntSlice(mat1[i], mat2[i]) {
			return false
		}
	}
	return true
}
