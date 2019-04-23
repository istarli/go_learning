package tree

import(
    "fmt"
    "strconv"
    "mymath"
)

func init() {
    fmt.Println("tree.go")
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func CreateTree(data []string) *TreeNode {
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
        begin,end = end,mymath.Min_int(end+num*2,len(data))
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