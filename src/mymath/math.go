package mymath
/*
	Simple math operation
*/

func Max_int(x,y int) int {
	if x < y {
		return y
	}
	return x
}

func Min_int(x,y int) int {
	if x < y {
		return x
	}
	return y
}