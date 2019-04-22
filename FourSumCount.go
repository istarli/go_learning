// Leetcode 454. 4Sum II
// https://leetcode-cn.com/problems/4sum-ii/submissions/

package main
import "fmt"

// Solution : hash-map
func fourSumCount(A []int, B []int, C []int, D []int) int {
    table := map[int]int{}
    for _,x := range A {
        for _,y := range B{
            table[x+y]++
        }
    }
    res := 0
    for _,x := range C {
        for _,y := range D {
            res += table[0-x-y]
        }
    }
    return res
}

// Local test
func main(){
    A := []int{1,2}
    B := []int{-2,-1}
    C := []int{-1,2}
    D := []int{0,2}
    fmt.Println(fourSumCount(A,B,C,D))
}