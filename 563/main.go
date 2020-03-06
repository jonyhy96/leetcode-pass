package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var result = 0
	traverse(root, &result)
	return result
}

func traverse(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left, result)
	right := traverse(root.Right, result)
	*result += abs(left, right)
	return left + right + root.Val
}

func abs(left, right int) int {
	if left < right {
		return right - left
	}
	return left - right
}

func main() {
	var left = TreeNode{Val: 2}
	var right = TreeNode{Val: 3}

	var root = &TreeNode{
		Val:   1,
		Left:  &left,
		Right: &right,
	}

	fmt.Println(findTilt(root))
}
