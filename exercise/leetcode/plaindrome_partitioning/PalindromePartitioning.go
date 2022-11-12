// Leetcode 131. Palindrome Partitioning
// https://leetcode-cn.com/problems/palindrome-partitioning/submissions/

package main
import "fmt"

// Solution1 : dfs
func helpFunc(i int, s string, v []string, res *[][]string){
    if len(s) == i {
        if check(v[len(v)-1]) {
            tmp := make([]string,len(v),len(v))
            copy(tmp,v)
            *res = append(*res, tmp)
        }
        return
    }

    if check(v[len(v)-1]) {
        v = append(v, s[i:i+1])
        helpFunc(i+1, s, v, res)
        v = v[:len(v)-1]
    }
    
    back := v[len(v)-1]
    v[len(v)-1] = back + s[i:i+1]
    helpFunc(i+1, s, v, res)
    v[len(v)-1] = back
}

// -->Optimize : save index in slice, no longer use copy
// TODO 
// func helpFunc2(){

// }

// Solution2 : bfs
// TODO
// func partition(s string)[][]string{

// }

func check(s string) bool {
    if "" == s {
        return false
    }
    for i,j := 0,len(s)-1; i < j; i,j = i+1,j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func partition(s string) [][]string {
    res := [][]string{}
    helpFunc(0, s, []string{""}, &res)
    return res
}

// Local test
func main(){
	s := "aab"
	fmt.Println(partition(s))
}