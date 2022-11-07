// Leetcode 152. Maximum Product Subarray
// https://leetcode-cn.com/problems/maximum-product-subarray/
package main
import "fmt"

/* Solution1 : 
	1. Split into subarrays by '0'
	2. For each subarray, check the num of negtive numbers and find the max
	3. Find the max one for each subarry's return-value. 
*/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func max1(a int, args ...int) int {
    res := a
    for _,x := range args {
        if res < x { res = x }
    }
    return res
}

func mul(nums []int) int {
    res := 1
    for _,x := range nums { 
        res *= x 
    }
    return res
}

func helpFunc(nums []int) int {
    if 0 == len(nums) {
        return 0
    }else if 1 == len(nums) {
        return nums[0]
    }
    cnt,first,last := 0,-1,-1
    for i,x := range nums {
        if x < 0 {
            cnt++
            last = i
            if -1 == first { first = i }
        }
    }
    res := 1
    if 0 == cnt&1 { // even
        res = mul(nums);
    }else{ // odd
        res = max1(mul(nums[0:last]),mul(nums[first+1:]))
    }
    return res
}

func maxProduct1(nums []int) int {
    res,begin := INT_MIN,0
    for i,x := range nums {
        if 0 == x {
            res = max1(res,0,helpFunc(nums[begin:i]))
            begin = i+1
        }
    }
    return max1(res,helpFunc(nums[begin:]))
}

/*
	Solution2 : dp, save min and max for every step in a for-loop
*/
func compare(comp func(int,int)bool, a int, args ...int) int {
    res := a
    for _,x := range args {
        if comp(x,res) { res = x }
    }
    return res
}

func max2(a int, args ...int) int {
    _max := func(a int,b int) bool {
        if a > b { 
            return true 
        }
        return false
    }
    return compare(_max,a,args...)
}

func min2(a int, args ...int) int {
    _min := func(a int,b int) bool {
        if a < b { 
            return true 
        }
        return false
    }
    return compare(_min,a,args...)
}

func maxProduct2(nums []int) int {
    if 0 == len(nums) {
        return 0
    }
    res,_max,_min := nums[0],1,1
    for _,x := range nums {
        val1,val2 := _max*x, _min*x
        _max,_min = max2(val1, val2, x), min2(val1,val2,x)
        res = max2(res,_max)
    }
    return res
}

// Local test
func main() {
	nums := []int{-1,-2,-3,0}
	fmt.Println(maxProduct2(nums))
}