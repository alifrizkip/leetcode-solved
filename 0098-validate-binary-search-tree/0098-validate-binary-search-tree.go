/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func isValidBST(root *TreeNode) bool {
    var left = pow(2, 32) * -1
    var right = pow(2, 32) - 1

    return check(root, left, right)
}

func check(node *TreeNode, left, right int) bool {
    if node == nil {
        return true
    }

    if !(left < node.Val && node.Val < right) {
        return false
    }

    return check(node.Left, left, node.Val) && check(node.Right, node.Val, right)
}

func pow(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}