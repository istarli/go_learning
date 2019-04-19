package main
import "fmt"
// Leetcode 207. Course Schedule 
// https://leetcode-cn.com/problems/course-schedule/

// Solution1 : dfs
func dfs(v int, matrix [][]int, flag []bool, safe []bool) {
    if safe[v] || flag[v] {
        return
    }
    flag[v] = true
    for _,x := range matrix[v] {
        dfs(x,matrix,flag,safe)
        if !safe[x] {return}
    }
    flag[v] = false
    safe[v] = true
}

func __canFinish1(num int, pre [][]int) bool {
    matrix := make([][]int,num)
    for _,link := range pre {
        v1,v2 := link[0],link[1]
        matrix[v1] = append(matrix[v1],v2)
    }
    flag := make([]bool,num)
    safe := make([]bool,num)
    for i := 0; i < num; i++ {
        dfs(i,matrix,flag,safe)
        if !safe[i] {return false}
    }
    return true
}

// Solution2 : bfs
func __canFinish2(num int, pre [][]int) bool {
    matrix := make([][]int,num)
    indegree := make([]int,num)
    for _,link := range(pre) {
        v1,v2 := link[0],link[1]
        matrix[v1] = append(matrix[v1],v2)
        indegree[v2]++;
    }
    q := make([]int,0)
    for i,x := range(indegree) {
        if 0 == x { q = append(q,i) }
    }
    cnt := 0
    for 0 != len(q) {
        v := q[0]
        q = q[1:]
        cnt++
        for _,x := range(matrix[v]) {
            indegree[x]--
            if 0 == indegree[x] {
                q = append(q,x)
            }
        }
    }
    return cnt == num;
}

// Local Test
func canFinish(num int, pre [][]int) bool {
	return __canFinish1(num,pre)
	// return __canFinish2(num,pre)
}

func main(){
	fmt.Println(canFinish(2,[][]int{{1,0}}))
	fmt.Println(canFinish(2,[][]int{{1,0},{0,1}}))
}