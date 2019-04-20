// Leetcode 230. Kth Smallest Element in a BST
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

package main
import (
    "fmt"
    "strconv"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// Solution : InOrder
func helpFunc(root *TreeNode, k int, cnt *int, flag *bool) (int,bool) {
    if nil == root {
        *flag = true
        return 0,false 
    }
    if val,ok := helpFunc(root.Left,k,cnt,flag); ok { 
        return val,ok 
    } else if *flag {
        if *cnt++; *cnt == k {
            return root.Val,true
        }
    }
    return helpFunc(root.Right,k,cnt,flag)
}

func kthSmallest(root *TreeNode, k int) int {
    cnt,flag := 0,false
    val,_ := helpFunc(root,k,&cnt,&flag)
    return val
}


// Local test
func Min(x,y int) int {
    if x < y { 
        return x 
    }
    return y
}

func createTree(data []string) *TreeNode {
    if 0 == len(data) || "null" == data[0] {
        return nil
    }
    val,_ := strconv.Atoi(data[0])
    root := TreeNode{val,nil,nil} 
    tree := [][]*TreeNode{{&root}}
    begin,end,depth := 1,3,1 // [1,3)
    for begin < len(data) {
        floor := data[begin:end]
        tree = append(tree,[]*TreeNode{})
        num := 0
        for i,s := range floor {
            if "null" != s {
                num++
                val,_ := strconv.Atoi(s)
                treeNode := TreeNode{val,nil,nil}
                tree[depth] = append(tree[depth],&treeNode)
                if 0 == i&1 {
                    tree[depth-1][i/2].Left = &treeNode
                }else{
                    tree[depth-1][i/2].Right = &treeNode
                }
            }
        }
        begin,end = end,Min(end+num*2,len(data))
        depth++
    }
    return &root
}

func PreOrder(root *TreeNode) {
    if nil != root {
        fmt.Print(root.Val," ")
        PreOrder(root.Left)
        PreOrder(root.Right)
    }
}

func InOrder(root *TreeNode) {
    if nil != root {
        InOrder(root.Left)
        fmt.Print(root.Val," ")
        InOrder(root.Right)
    }
}

func main(){
    root := createTree([]string{
        "5",
        "0","8",
        "null","3","6","10",
        "2","4","null","7","9"})
    // test for createTree
    PreOrder(root)
    fmt.Println()
    InOrder(root)
    fmt.Println()
    // test for KthSmallest
    fmt.Println(kthSmallest(root,4))
}