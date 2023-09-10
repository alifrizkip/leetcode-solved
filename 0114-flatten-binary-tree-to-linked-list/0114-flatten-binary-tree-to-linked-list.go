/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
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

    if len(nodes) <= 1 {
        return
    }

    nodes = nodes[1:]
    for len(nodes) > 0 {
        n := nodes[0]
        nodes = nodes[1:]

        root.Left = nil
        root.Right = n
        root = root.Right
    }
}