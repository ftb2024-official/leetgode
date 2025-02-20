/*
Given the `root` of a binary search tree and an integer `k`, return `true` if there exist two elements in the BST such that their sum is equal to
`k`, or `false` otherwise.

Example 1:
	Input: root = [5,3,6,2,4,null,7], k = 9
	Output: true

Example 2:
	Input: root = [5,3,6,2,4,null,7], k = 28
	Output: false
*/

package arrays_and_hashing

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, k int, visited map[int]bool) bool {
	if root == nil {
		return false
	}

	if visited[k-root.Value] {
		return true
	}
	visited[root.Value] = true

	return dfs(root.Left, k, visited) || dfs(root.Right, k, visited)
}

func bfs(root *TreeNode, k int, visited map[int]bool) bool {
	if root == nil {
		return false
	}

	if visited[root.Value] {
		return true
	}
	visited[k-root.Value] = true

	return bfs(root.Left, k, visited) || bfs(root.Right, k, visited)
}

func FindTarget(root *TreeNode, k int) bool {
	return bfs(root, k, map[int]bool{})
}
