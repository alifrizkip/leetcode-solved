/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    var res int
    q := []*TreeNode{root}
    nextQ := []*TreeNode{}

    for len(q) > 0{
        n := q[0]
        q = q[1:]

        res += n.Val

        if n.Left != nil {
            nextQ = append(nextQ, n.Left)
        }
        if n.Right != nil {
            nextQ = append(nextQ, n.Right)
        }

        if len(q) == 0 && len(nextQ) > 0 {
            res = 0
            q = nextQ
            nextQ = []*TreeNode{}
        }
    }

    return res
}