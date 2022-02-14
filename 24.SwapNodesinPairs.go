/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head}
    dummy := &ListNode{0, nil}
    curr := dummy
    dummy.Next = head
    
    for curr.Next != nil && curr.Next.Next != nil {
        p1, p2 := curr.Next, curr.Next.Next
        
        curr.Next = p2
        p1.Next = p2.Next
        p2.Next = p1
        
        curr = p1
    }
    
    return dummy.Next
    
}