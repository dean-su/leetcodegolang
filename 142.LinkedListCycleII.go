/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    f, s := head, head
    cycled := false
    for f != nil && f.Next != nil {
        f, s = f.Next.Next, s.Next
        if f == s {
            cycled = true
            break;
        }
    }
    if cycled == false {return nil}
    
    s = head
    for s != f {
        s, f = s.Next, f.Next
    }
    return s
}