/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 currSum
 visit list
 if node = 0 && currSum != 0
    add to res
*/
func mergeNodes(head *ListNode) *ListNode {
    res := &ListNode{
        Next: &ListNode{},
    }
    node := res.Next

    var currSum int
    for head != nil {
        currSum += head.Val

        if head.Val == 0 && currSum > 0 {
            node.Val = currSum
            node.Next = &ListNode{}
            
            // end of head
            // shoud be end of node
            if head.Next == nil {
                node.Next = nil
            }

            node = node.Next
            currSum = 0
        }

        head = head.Next
    }

    return res.Next
}