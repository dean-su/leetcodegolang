/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
func hasCycle(head *ListNode) bool {
    if head == nil {return false}
    visited := make(map[*ListNode]bool)
    
    for head != nil {
        if visited[head] == true {return true}
        visited[head] = true
        head = head.Next
    }
    return false
}
*/

 func hasCycle(head *ListNode) bool {
    if head == nil {return false}
    f, s := head, head
    for f != nil {
        if f.Next == nil {return false}
        f = f.Next.Next
        s = s.Next
        if f == s {return true}
    }
    return false
}