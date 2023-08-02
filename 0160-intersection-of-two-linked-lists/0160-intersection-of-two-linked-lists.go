/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var mapA = map[*ListNode]bool{}
    var mapB = map[*ListNode]bool{}
    var nodeA, nodeB = headA, headB

    for nodeA != nil || nodeB != nil {
        mapA[nodeA] = true
        mapB[nodeB] = true
        
        if mapB[nodeA] {
            return nodeA
        }
        
        if mapA[nodeB] {
            return nodeB
        }
        
        if nodeA != nil {
            nodeA = nodeA.Next    
        }
        if nodeB != nil {
            nodeB = nodeB.Next
        }
    }
    
    return nil
}