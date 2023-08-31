/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
   1 2 4
   l
   1 3 4
   r
 0
 */
func mergeTwoLists(l *ListNode, r *ListNode) *ListNode {
    var head = new(ListNode)
    var curr = head

    for l != nil && r != nil {
        if l.Val < r.Val {
            curr.Next = l
            l = l.Next
        } else {
            curr.Next = r
            r = r.Next
        }
        curr = curr.Next
    }

    if l == nil {
        curr.Next = r
    } else {
        curr.Next = l
    }

    return head.Next
}