/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {return [][]int{}}
    res := [][]int{}
    curr := []*TreeNode{}
    curr = append(curr, root)
    
    for len(curr) != 0 {
        next := []*TreeNode{}
        currVal := []int{}
        for _, node :=  range curr {
            currVal = append(currVal, node.Val)
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
            
        }
        curr = next
        res = append(res, currVal)
        
    }
    
    for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
    
    return res
}