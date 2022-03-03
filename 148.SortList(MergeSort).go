/*
148. Sort List
Medium

Given the head of a linked list, return the list after sorting it in ascending order.

 

Example 1:


Input: head = [4,2,1,3]
Output: [1,2,3,4]
Example 2:


Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
Example 3:

Input: head = []
Output: []
 

Constraints:

The number of nodes in the list is in the range [0, 5 * 104].
-105 <= Node.val <= 105
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head;}
    slow := head
    fast := head.Next
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    mid := slow.Next
    slow.Next = nil
    return merge(sortList(head), sortList(mid))
}

func merge(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{0,nil}
    tail := dummy
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            l1.Val, l2.Val = l2.Val, l1.Val
            l1.Next, l2.Next = l2.Next, l1.Next
        }
        
        tail.Next = l1
        l1 = l1.Next
        tail = tail.Next  
    } 
    if l1 != nil {
        tail.Next = l1
    }else {
        tail.Next = l2
    }
    
    return dummy.Next
    
}