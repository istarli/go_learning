package data_structure

func InsertSortCore(seq []int, d int) {
	// if seq is nil, len(seq) will get zero
	for i := d; i < len(seq); i += d {
		if seq[i] < seq[i-d] {
			tmp := seq[i]
			j := i - d
			for ; j >= 0 && seq[j] > tmp; j -= d {
				seq[j+d] = seq[j]
			}
			seq[j+d] = tmp
		}
	}
}

func InsertSort(seq []int) {
	InsertSortCore(seq, 1)
}

func ShellSort(seq []int) {
	for d := len(seq) / 2; d > 0; d /= 2 {
		InsertSortCore(seq, d)
	}
}

func BubbleSort(seq []int) {
	pos := len(seq) - 1
	for pos > 0 {
		bound := pos
		pos = 0
		for j := 0; j < bound; j++ {
			if seq[j] > seq[j+1] {
				seq[j], seq[j+1] = seq[j+1], seq[j]
				pos = j
			}
		}
	}
}

func QuickSort(seq []int) {
	if len(seq) > 0 {
		k := Partion(seq)
		QuickSort(seq[0:k])
		QuickSort(seq[k+1:])
	}
}

func Partion(seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	pivot, i, j := seq[0], 0, len(seq)-1
	for i < j {
		for i < j && seq[j] >= pivot {
			j--
		}
		seq[i] = seq[j]
		for i < j && seq[i] <= pivot {
			i++
		}
		seq[j] = seq[i]
	}
	seq[i] = pivot
	return i
}

func SelectSort(seq []int) {
	for i := 0; i < len(seq); i++ {
		minIdx := i
		for j := i + 1; j < len(seq); j++ {
			if seq[j] < seq[minIdx] {
				minIdx = j
			}
		}
		seq[i], seq[minIdx] = seq[minIdx], seq[i]
	}
}

func HeapSort(seq []int) {
	n := len(seq)
	// create heap
	for i := n / 2; i >= 0; i-- {
		Sift(seq, i)
	}
	// pop top
	for i := n - 1; i > 0; i-- {
		seq[0], seq[i] = seq[i], seq[0]
		Sift(seq[:i], 0)
	}
}

func Sift(seq []int, k int) {
	if k < 0 {
		return
	}
	i, j := k, 2*k+1
	for j < len(seq) {
		if j+1 < len(seq) && seq[j] < seq[j+1] {
			j++
		}
		if seq[i] >= seq[j] {
			break
		}
		seq[i], seq[j] = seq[j], seq[i]
		i, j = j, 2*j+1
	}
}

func Merge(seq1 []int, seq2 []int) {

}
