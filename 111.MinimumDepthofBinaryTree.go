/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {return 0}
    if root.Left == nil && root.Right == nil {return 1}
    l, r := math.MaxInt, math.MaxInt
    if root.Left != nil {
        l = minDepth(root.Left)
    }
    if root.Right != nil {
        r = minDepth(root.Right)
    }
    
    return min(l, r) + 1
}

func min(a, b int) int {
    if a<b {
        return a
    }else {
        return b
    }
}