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

/*
	Time Complexity : O(n)~O(n^2), O(n^2)
	Space Complexity : O(1)
	Stable.
*/
func InsertSort(seq []int) {
	InsertSortCore(seq, 1)
}

/*
	Time Complexity : O(n*log2(n))~O(n^2), O(n^1.3)
	Space Complexity : O(1)
	Not Stable.
*/
func ShellSort(seq []int) {
	for d := len(seq) / 2; d > 0; d /= 2 {
		InsertSortCore(seq, d)
	}
}

/*
	Time Complexity : O(n)~O(n^2), O(n^2)
	Space Complexity : O(1)
	Stable.
*/
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

/*
	Time Complexity : O(n*log2(n))~O(n^2), O(n*log2(2))
	Space Complexity : O(1)
	Not Stable.
*/ 
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

/*
	Time Complexity : O(n^2)
	Space Complexity : O(1)
	Stable.
	The least times to move elements.
*/ 
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

/*
	Time Complexity : O(n*log2(n))
	Space Complexity : O(1)
	Stable.
*/ 
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

/*
	Time Complexity : O(n*log2(n))
	Space Complexity : O(n)
	Stable.
*/ 
func MergeSort(seq []int) {
	if len(seq) > 1 {
		mid := len(seq) / 2
		MergeSort(seq[0:mid])
		MergeSort(seq[mid:])
		Merge(seq[0:mid], seq[mid:])
	}
}

func MergeSortNR(seq []int) {
	d, n := 1, len(seq)
	for d < n {
		for i := 0; i < n; i += 2 * d {
			if end, mid := i+2*d, i+d; end <= n {
				Merge(seq[i:mid], seq[mid:end])
			} else if mid <= n {
				Merge(seq[i:mid], seq[mid:])
			}
		}
		d *= 2
	}
}

func Merge(seq1 []int, seq2 []int) {
	n1, n2 := len(seq1), len(seq2)
	if n1 >= 0 && n2 >= 0 {
		i, j, k := 0, 0, 0
		tmp := make([]int, n1+n2)
		for i < n1 && j < n2 {
			if seq1[i] < seq2[j] {
				tmp[k] = seq1[i]
				i, k = i+1, k+1
			} else {
				tmp[k] = seq2[j]
				j, k = j+1, k+1
			}
		}
		for i < n1 {
			tmp[k] = seq1[i]
			i, k = i+1, k+1
		}
		for j < n2 {
			tmp[k] = seq2[j]
			j, k = j+1, k+1
		}
		for m := 0; m < n1; m++ {
			seq1[m] = tmp[m]
		}
		for m := 0; m < n2; m++ {
			seq2[m] = tmp[n1+m]
		}
	}
}
