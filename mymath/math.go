package mymath
/*
	Simple math operation
*/

import "fmt"

func init() {
	fmt.Println("math.go")
}

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