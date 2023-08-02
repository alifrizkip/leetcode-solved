/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    var slow, fast = head, head

    /*
        Tortoise & Hare algorithm
        Tortoise move slow by 1
        Hare move fast by 2
        Cycle detected if Hare overlap Tortoise
    */
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if slow == fast {
            return true
        }
    }

    return false
}