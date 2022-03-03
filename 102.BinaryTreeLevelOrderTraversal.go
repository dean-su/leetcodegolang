/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    dfs(root, 0, &res)
    return res
}*/

func dfs(root *TreeNode, depth int, res *[][]int) {
    if root == nil {return}
    for len(*res) <= depth {
        *res = append(*res, []int{})
    }
    
    (*res)[depth] = append((*res)[depth], root.Val)
    dfs(root.Left, depth+1, res)
    dfs(root.Right, depth+1, res)
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {return [][]int{}}
    res := [][]int{}
    curr := []*TreeNode{}
    curr = append(curr, root)
    
    for len(curr) != 0 {

        next := []*TreeNode{}
        currVal := make([]int, 0)
        for _, node := range curr {
            currVal = append(currVal, node.Val)
            if node.Left != nil { 
                next = append(next, node.Left) 
            }
            if node.Right != nil { 
                next = append(next, node.Right)
            }
            
        }
        res = append(res, currVal)
        curr = next
    }
    return res
}