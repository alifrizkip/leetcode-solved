/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    var res int
    queue := []*TreeNode{root}
    nextQueue := []*TreeNode{}

    for len(queue) > 0{
        root = queue[0]
        queue = queue[1:]

        if root == nil {
            continue
        }

        if root.Left != nil {
            nextQueue = append(nextQueue, root.Left)
        }

        if root.Right != nil {
            nextQueue = append(nextQueue, root.Right)
        }

        if len(queue) == 0 {
            res++
            queue = nextQueue
            nextQueue = []*TreeNode{}
        }
    }

    return res
}