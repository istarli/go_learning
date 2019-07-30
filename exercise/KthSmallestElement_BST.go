// Leetcode 230. Kth Smallest Element in a BST
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

package main

import (
	"fmt"

	"github.com/istarli/go_learning/tree"
)

type TreeNode = tree.TreeNode

// Solution : InOrder
func helpFunc(root *TreeNode, k int, cnt *int, flag *bool) (int, bool) {
	if nil == root {
		*flag = true
		return 0, false
	}
	if val, ok := helpFunc(root.Left, k, cnt, flag); ok {
		return val, ok
	} else if *flag {
		if *cnt++; *cnt == k {
			return root.Val, true
		}
	}
	return helpFunc(root.Right, k, cnt, flag)
}

func kthSmallest(root *TreeNode, k int) int {
	cnt, flag := 0, false
	val, _ := helpFunc(root, k, &cnt, &flag)
	return val
}

// Local test
func main() {
	root := tree.CreateTree([]string{
		"5",
		"0", "8",
		"null", "3", "6", "10",
		"2", "4", "null", "7", "9"})
	// test for createTree
	tree.PreOrder(root)
	fmt.Println()
	tree.InOrder(root)
	fmt.Println()
	// test for KthSmallest
	fmt.Println(kthSmallest(root, 4))
}
