/*
Given the roots of two binary trees `p` and `q`, write a funtion to check if they are the same or not. Two binary trees are considered the same
if they are structurally identical, and the nodes have the same value.

Example 1:
	Input: p = [1,2,3], q = [1,2,3]
	Output: true

Example 2:
	Input: p = [1,2], q = [1,null,2]
	Output: false

Example 3:
	Input: p = [1,2,1], q = [1,1,2]
	Output: false
*/

package yandex

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right
}
