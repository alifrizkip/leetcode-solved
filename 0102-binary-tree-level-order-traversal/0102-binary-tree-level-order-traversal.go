/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
 create queue
 add root to queue
 loop until len(queue):
    init row list
    dequeue item
    add item.left,item.right to queue
    add dequeued item to row list
add row list to res
 */
func levelOrder(root *TreeNode) (res [][]int) {
    if root == nil {
        return res
    }

    q := []*TreeNode{root}
    for len(q) > 0 {
        qL := len(q)
        row := []int{}

        for i := 0; i < qL; i++ {
            item := q[0]
            row = append(row, item.Val)

            if item.Left != nil {
                q = append(q, item.Left)
            }

            if item.Right != nil {
                q = append(q, item.Right)
            }

            q = q[1:]
        }

        res = append(res, row)
    }

    return res
}