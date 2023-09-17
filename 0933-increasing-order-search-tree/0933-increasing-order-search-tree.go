/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    arr := visit(root, []int{})

    dummy := &TreeNode{
        Right: &TreeNode{},
    }
    node := dummy.Right

    for len(arr) > 0 {
        val := arr[0]
        arr = arr[1:]

        node.Val = val

        if len(arr) > 0 {
            node.Right = &TreeNode{}
        }

        node = node.Right
    }

    return dummy.Right
}

func visit(node *TreeNode, arr []int) []int {
    if node == nil {
        return arr
    }

    if node.Left != nil {
        arr = visit(node.Left, arr)
    }

    arr = append(arr, node.Val)

    if node.Right != nil {
        arr = visit(node.Right, arr)
    }

    return arr
}