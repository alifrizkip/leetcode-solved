/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    var revHead *ListNode
    for head != nil {
        tmpNext := head.Next

        head.Next = revHead
        revHead = head

        head = tmpNext
    }

    return revHead
}

func reverseListArray(head *ListNode) *ListNode {
    var arr []int
    ln := head
    for ln != nil {
        arr = append(arr, ln.Val)
        ln = ln.Next
    }

    var node *ListNode
    for _, num := range arr {
        node = &ListNode{
            Val: num,
            Next: node,
        }
    }

    return node
}