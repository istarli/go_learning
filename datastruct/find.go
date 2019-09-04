package datastruct

func BinarySearch(seq []int, target int) int {
	low, high := 0, len(seq)
	for low < high {
		mid := low + (high-low)/2
		if seq[mid] > target {
			high = mid
		} else if seq[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func LowerBound(seq []int, target int) int {
	low, high := 0, len(seq)
	for low < high {
		mid := low + (high-low)/2
		if seq[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func UpperBound(seq []int, target int) int {
	low, high := -1, len(seq)-1
	for low < high {
		mid := high - (high-low)/2
		if seq[mid] <= target {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return high
}
