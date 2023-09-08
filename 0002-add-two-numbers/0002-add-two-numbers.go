/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l, r := l1, l2

    res := &ListNode{
        Next: &ListNode{},
    }
    node := res.Next

    for (l != nil || r != nil) && node != nil {
        node.Next = &ListNode{}

        var lv, rv int
        if l != nil {
            lv = l.Val
            l = l.Next
        }
        if r != nil {
            rv = r.Val
            r = r.Next
        }

        sum := lv + rv + node.Val
        if sum >= 10 {
            node.Next.Val = 1
            sum = sum - 10
        } else {
            node.Next.Val = 0
        }

        node.Val = sum

        if l == nil && r == nil && node.Next.Val == 0 {
            node.Next = nil
        }

        node = node.Next
    }

    return res.Next
}