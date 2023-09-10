/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Space: O(1)
func flatten(root *TreeNode) {
    var prev *TreeNode
    var visit func(*TreeNode)
    visit = func(node *TreeNode) {
        if node == nil {
            return
        }

        visit(node.Right)
        visit(node.Left)

        node.Right = prev
        node.Left = nil
        prev = node
    }
    visit(root)
}

// Space: O(n)
func flattenV1(root *TreeNode)  {
    nodes := []*TreeNode{}

    var visit func(*TreeNode)
    visit = func(node *TreeNode) {
        if node == nil {
            return
        }

        nodes = append(nodes, node)
        visit(node.Left)
        visit(node.Right)
    }
    visit(root)

    for i := 0; i < len(nodes)-1; i++ {
        nodes[i].Left = nil
        nodes[i].Right = nodes[i+1]
    }
}