/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    res := []float64{}

    q := []*TreeNode{root}
    nextQ := []*TreeNode{}

    var tmpSum, tmpCount int
    for len(q) > 0 {
        n := q[0]
        q = q[1:]

        tmpSum += n.Val
        tmpCount++

        if n.Left != nil {
            nextQ = append(nextQ, n.Left)
        }
        if n.Right != nil {
            nextQ = append(nextQ, n.Right)
        }

        if len(q) == 0 {
            q = nextQ
            nextQ = []*TreeNode{}

            var avg float64 = float64(tmpSum) / float64(tmpCount)
            res = append(res, avg)

            tmpSum, tmpCount = 0, 0
        }
    }

    return res
}