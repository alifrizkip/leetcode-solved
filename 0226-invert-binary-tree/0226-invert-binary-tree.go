/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    dfs(root)
    return root
}

func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    node.Left, node.Right = node.Right, node.Left
    dfs(node.Left)
    dfs(node.Right)
}