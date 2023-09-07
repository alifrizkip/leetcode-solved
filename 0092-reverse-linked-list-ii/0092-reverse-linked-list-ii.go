/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    arrVal :=[]int{}

    node := head
    i := 0
    for node != nil {
        i++

        if i >= left && i <= right {
            arrVal = append(arrVal, node.Val)
        }

        node = node.Next
    }

    res := &ListNode{
        Next: head,
    }
    node = res.Next
    i = 0
    for node != nil {
        i++

        if i >= left && i <= right {
            node.Val = arrVal[len(arrVal)-1]
            arrVal = arrVal[:len(arrVal)-1]
        }

        node = node.Next
    }

    return res.Next
}