/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var listLen int

    res, node := head, head
    for node != nil {
        listLen++
        node = node.Next
    }

    var start int
    mid := listLen / 2
    for start < mid {
        start++
        res = res.Next
    }

    return res
}