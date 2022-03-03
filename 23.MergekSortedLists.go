/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }
    
    for len(lists) > 1 {
        // pop 2 lists
        l1 := lists[0]
        l2 := lists[1]
        lists = lists[2:]
        
        merged := mergeTwoLists(l1, l2)
        lists = append(lists, merged)
    }
    
    return lists[0]
}


func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }
    
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    }
}

 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }
    
    for len(lists) > 1 {
        // pop 2 lists
        l1 := lists[0]
        l2 := lists[1]
        lists = lists[2:]
        
        merged := mergeTwoLists(l1, l2)
        lists = append(lists, merged)
    }
    
    return lists[0]
}


func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }
    
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLists(l2.Next, l1)
        return l2
    }
}