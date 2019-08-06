package data_structure

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
