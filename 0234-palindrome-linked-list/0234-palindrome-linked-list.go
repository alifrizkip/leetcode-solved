/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
    1. convert to array
*/
func isPalindrome(head *ListNode) bool {
    arr := []int{}
    node := head
    for node != nil {
        arr = append(arr, node.Val)
        node = node.Next
    }
    
    for i := 0; i < len(arr); i++ {
        if arr[i] != arr[len(arr)-1-i] {
            return false
        }
    }
    
    return true
}