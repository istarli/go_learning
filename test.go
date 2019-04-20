package main

import(
	"fmt"
)

func test() *[]int {
	slice := []int{4,5,6}
	return &slice
}

func main(){
	fmt.Println(*test())
}