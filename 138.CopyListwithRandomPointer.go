/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func copyRandomList(head *Node) *Node {
    if head == nil {return head}
    
    m := make(map[*Node]*Node)
    curr := head
    
    for curr != nil {
        dummy := new(Node)
        dummy.Val = curr.Val
        m[curr] = dummy
        curr = curr.Next
    }
    
    curr = head
    for curr != nil {
        m[curr].Next = m[curr.Next]
        m[curr].Random = m[curr.Random]
        curr = curr.Next
    }
    
    return m[head]
}